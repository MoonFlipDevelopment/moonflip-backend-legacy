package main

import (
	"bytes"
	"context"
	cryptoRand "crypto/rand"
	"fmt"
	"log"
	"math"
	"math/big"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/MoonFlipDevelopment/moonflip/contracts/moonflip"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
)

const gwei = 1000000000

type Chain struct {
	PrettyName  string
	Symbol      string
	ChainID     *big.Int
	RPCProvider string
	RPCClient   *ethclient.Client

	GameContract common.Address

	MoonflipProgram *moonflip.Moonflip
}

func (chain *Chain) HandleTransaction(primaryAccount *keystore.Key) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		rpcClient := chain.RPCClient
		moonflipProgram := chain.MoonflipProgram
		chainId := chain.ChainID

		txnString := c.Params("transaction")

		// verify that this is a transaction
		if txnPattern.MatchString(txnString) {

			if !strings.HasPrefix(txnString, "0x") {
				txnString = "0x" + txnString
			}

			// check if we've already received a transcation
			var exists bool

			ExistsInDatabase(DatabaseTable, db, db.Table(DatabaseTable).Where("transaction_hash = ?", txnString)).Scan(&exists)

			if exists {
				//we have already processed this transaction. This is a POST request, not a GET, we'll have a separate GET functionality for looking up transactions
				log.Println("already processed transaction:", txnString)
				var gameResult GameResult

				db.Table(GameTable).Select("*").Where("commit_transaction = ? OR response_transaction = ?", txnString, txnString).Scan(&gameResult)

				if gameResult.CommitTxnHash == "" && gameResult.ResponseTxnHash == "" {
					return c.Status(http.StatusTooEarly).SendString("game transaction not fully processed")
				}

				return c.Status(http.StatusOK).JSON(gameResult)
			}

			txnHash := common.HexToHash(txnString)

			txn, isPending, err := rpcClient.TransactionByHash(context.Background(), txnHash)

			if isPending {
				log.Println(err)
				return c.Status(http.StatusTooEarly).SendString("transaction is pending")
			}

			if err != nil {
				log.Println(err)
				return c.Status(http.StatusBadRequest).SendString("transaction failed.")
			}

			if txn.To().String() == chain.GameContract.Hex() {
				receipt, err := rpcClient.TransactionReceipt(context.Background(), txnHash)

				if err != nil {
					log.Println(err)
					return c.Status(http.StatusInternalServerError).SendString(err.Error())
				}

				if len(receipt.Logs) != 1 {
					log.Printf("invalid log length: %d\n", len(receipt.Logs))
					return c.Status(http.StatusBadRequest).SendString(fmt.Sprintf("invalid log length: %d\n", len(receipt.Logs)))
				}

				guess, err := moonflipProgram.MoonflipFilterer.ParseGuessSubmitted(*receipt.Logs[0])

				if err != nil {
					log.Printf("invalid log format: %s\n", err.Error())
					return c.Status(http.StatusBadRequest).SendString(fmt.Sprintf("invalid log format: %s\n", err.Error()))
				}

				// The guess txn
				txn := &Transaction{
					From:            guess.Sender.String(),
					To:              guess.Raw.Address.String(),
					TxnHash:         guess.Raw.TxHash.String(),
					Value:           NewBigInt(guess.Commit.Value),
					Gas:             NewBigInt(new(big.Int).SetUint64(txn.Gas() * txn.GasPrice().Uint64())),
					TransactionType: Guess,
					CreatedAt:       time.Now(),
					Chain:           chain.Symbol,
				}

				// create the transaction first.
				if err = db.Create(txn).Error; err != nil {
					return c.Status(http.StatusBadRequest).SendString(err.Error())
				}

				buf := make([]byte, 20)

				cryptoRand.Read(buf)

				random, err := cryptoRand.Int(bytes.NewBuffer(buf), big.NewInt(math.MaxInt64))

				if err != nil {
					log.Println("failed to generate cryptographically secure random number...using P-RNG fallback.")

					rand.Seed(time.Now().UnixNano())
					random = big.NewInt(rand.Int63n(math.MaxInt64))
				}

				choice := Tails

				// the thing is, we already know if they'd win or not. So.
				// Just keep this logic, but submit on-chain for transparency too.
				// instead of if(won), check how they won, on chain.
				if new(big.Int).Mod(random, big.NewInt(2)).Int64() == 0 {
					choice = Heads
				}

				won := guess.Commit.Guess == uint8(choice)

				gasPriceEstimate, err := rpcClient.SuggestGasPrice(context.Background())

				if err != nil {
					log.Println("failed to estimate gas price", err)
					return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("failed to estimate gas price: %s", err.Error()))
				}

				gasTipEstimate, err := rpcClient.SuggestGasTipCap(context.Background())

				if err != nil {
					log.Println("failed to estimate gas tip", err)
					return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("failed to estimate gas tip: %s", err.Error()))
				}

				processTransaction, err := moonflipProgram.ProcessGuess(&bind.TransactOpts{
					From:      primaryAccount.Address,
					Signer:    SignerFunction(chainId),
					GasLimit:  0,
					GasFeeCap: gasPriceEstimate,
					GasTipCap: gasTipEstimate,
				}, guess.Sender, guess.Commit.Nonce, random)

				if err != nil {
					log.Printf("failed to submit transaction: %s\n", err.Error())
					return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("failed to submit transaction: %s\n", err.Error()))
				}

				value := guess.Commit.Value

				if won {
					value = value.Mul(value, big.NewInt(2)) // multiply the value by 2
				}

				gameResult := &GameResult{
					Player:          guess.Sender.String(),
					CommitTxnHash:   txnHash.String(),
					ResponseTxnHash: processTransaction.Hash().String(),
					Result:          won,
					GameId:          NewBigInt(guess.Commit.Nonce),
					Value:           NewBigInt(value),
					Chain:           chain.Symbol,
				}

				err = db.Table(GameTable).Create(gameResult).Error

				if err != nil {
					log.Println("failed to insert into the database:", err)
					return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("failed to insert into the database: %s\n", err.Error()))
				}

				transactions := []*Transaction{
					// The reward transaction
					{
						From:            primaryAccount.Address.String(),
						To:              processTransaction.To().String(),
						TxnHash:         processTransaction.Hash().String(),
						Value:           NewBigInt(value),
						Gas:             NewBigInt(new(big.Int).SetUint64(processTransaction.Gas() * processTransaction.GasPrice().Uint64())),
						TransactionType: Reward, // not really a reward. no guarantee
						CreatedAt:       time.Now(),
						Chain:           chain.Symbol,
					},
				}

				log.Println("submitted processed transaction:", processTransaction.Hash().String())
				// only insert the guess transaction if the processed transaction is actually submitted
				err = db.Table(DatabaseTable).Create(transactions[0]).Error

				if err != nil {
					log.Println("failed to insert into the database:", err)
					return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("failed to insert into the database: %s\n", err.Error()))
				}

				return c.Status(http.StatusOK).JSON(gameResult)
			}

			return c.Status(http.StatusBadRequest).SendString(fmt.Sprintf("invalid contract: %s\n", txn.To().String()))
		}

		return c.Status(http.StatusBadRequest).SendString(fmt.Sprintf("invalid txn hash: %s\n", txnString))
	}

}

func NewMainnetChain(prettyName, symbol, rpcProvider, contractAddress string, chainId *big.Int) *Chain {
	return NewChain(prettyName, symbol, rpcProvider, contractAddress, chainId, false)
}

func NewChain(prettyName, symbol, rpcProvider, contractAddress string, chainId *big.Int, testnet bool) *Chain {
	client, err := ethclient.Dial(rpcProvider)

	if err != nil {
		log.Panicln(err)
		return nil
	}

	var moonflipProgram *moonflip.Moonflip
	var address common.Address

	if contractAddress != "" {
		address = common.HexToAddress(contractAddress)

		contract, err := moonflip.NewMoonflip(address, client)

		if err != nil {
			log.Panicln(err)
			return nil
		}

		moonflipProgram = contract
	}

	return &Chain{
		PrettyName:  prettyName,
		Symbol:      symbol,
		RPCProvider: rpcProvider,
		RPCClient:   client,
		ChainID:     chainId,

		MoonflipProgram: moonflipProgram,
		GameContract:    address,
	}
}

var Chains []*Chain = []*Chain{
	NewMainnetChain("Moonriver", "MOVR", "https://rpc.moonriver.moonbeam.network", "0x7d5de78E0f2b73C3D6324D65A1a765dF59424713", big.NewInt(1285)),
	NewMainnetChain("Avalanche", "AVAX", "https://api.avax.network/ext/bc/C/rpc", "0x7d5de78E0f2b73C3D6324D65A1a765dF59424713", big.NewInt(43114)),
	NewChain("Goerli (Ethereum Testnet)", "GOETH", "https://goerli.infura.io/v3/", "", big.NewInt(5), true),
	NewChain("Moonbase Alpha (Moonbeam/Moonriver Testnet)", "DEV", "https://rpc.api.moonbase.moonbeam.network", "", big.NewInt(1287), true),
	NewChain("Avalanche Fuji", "tAVAX", "https://api.avax-test.network/ext/bc/C/rpc", "0x7d5de78E0f2b73C3D6324D65A1a765dF59424713", big.NewInt(43113), true),
}

var PathsToChainsLookups = map[string]*Chain{
	"movr":     Chains[0],
	"avax":     Chains[1],
	"goerli":   Chains[2],
	"moonbase": Chains[3],
	"fuji":     Chains[4],
}

func GetChain(path string) *Chain {
	return PathsToChainsLookups[path]
}
