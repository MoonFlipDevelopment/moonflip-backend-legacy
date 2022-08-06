package main

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/MoonFlipDevelopment/moonflip/contracts/moonflip"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

const (
	Guess  = 0
	Reward = 1
)

type BigInt struct {
	internal *big.Int
}

func NewBigInt(b *big.Int) *BigInt {
	return &BigInt{
		internal: b,
	}
}

func (i *BigInt) Scan(value interface{}) error {
	str, ok := value.(string)

	if !ok {

		b, ok := value.([]byte)

		if ok {
			str = string(b)
		} else {
			return errors.New(fmt.Sprint("Failed to unmarshal BigInt value:", value))
		}
	}

	bi, ok := new(big.Int).SetString(str, 10)

	if !ok {
		return errors.New(fmt.Sprint("Failed to set big.Int string", str))
	}

	i.internal = bi
	return nil
}

func (i BigInt) Value() (driver.Value, error) {
	if i.internal == nil {
		return big.NewInt(0).String(), nil
	}
	return i.internal.String(), nil
}

func (i BigInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.internal.String())
}

type User struct {
	// These are optional
	DiscordID  string `gorm:"type:varchar(255);column:discord_id;" json:"discordId"`
	DiscordTag string `gorm:"type:varchar(255);column:discord_tag;" json:"discordTag"`
	Address    string `gorm:"type:varchar(255);column:address;" json:"address"`
}

func (*User) TableName() string {
	return "users"
}

type Session struct {
	// The address for this session
	Address common.Address
	// The session ID
	SessionID string
}

type PendingTransactions struct {
	Txns    []*Transaction
	Event   *moonflip.MoonflipGuessSubmitted
	TxnHash string
}

func (pending *PendingTransactions) IsCompleted(client *ethclient.Client) (completed bool, failedReason error) {
	transaction, isPending, err := client.TransactionByHash(context.Background(), common.HexToHash(pending.TxnHash))

	if isPending {
		return false, ErrPendingTransaction
	}

	if err != nil {
		return false, err
	}

	receipt, err := client.TransactionReceipt(context.Background(), transaction.Hash())

	if err != nil {
		return false, err
	}

	if receipt.Status != 0x1 {
		if receipt.GasUsed == receipt.CumulativeGasUsed {
			return false, ErrOutOfGas
		}
		return false, ErrFailedTransaction
	}

	return receipt.Status == 0x1, nil
}

func (pending *PendingTransactions) InsertIntoDB(db *gorm.DB) error {
	err := db.Table(DatabaseTable).CreateInBatches(pending.Txns, len(pending.Txns)).Error

	if err != nil {
		log.Println("failed to insert into the database:", err)
		return err
	}
	return nil
}

// This is a processed transaction
type Transaction struct {
	// The transaction hash
	TxnHash string `gorm:"type:varchar(255);unique;primaryKey;column:transaction_hash;" json:"txnHash"`

	// The from address
	From string `gorm:"type:varchar(255);column:from;" json:"from"`

	//
	To string `gorm:"type:varchar(255);column:to;" json:"to"`
	// The true value of the txn (with gas + fee)
	Value *BigInt `gorm:"type:varchar(255);column:value;" json:"value"`

	// How much gas this transaction used
	Gas *BigInt `gorm:"type:varchar(255);column:gas;" json:"gas"`

	// The type of this transaction
	TransactionType int `gorm:"type:int;column_transaction_type" json:"txnType"`

	// The time this transaction was created at
	CreatedAt time.Time `gorm:"type:time;column:created_at;" json:"createdAt"`

	Chain string `gorm:"type:varchar(255);column:chain;" json:"chain"`
}

func (*Transaction) TableName() string {
	return "transactions"
}

type GameResult struct {
	Player          string  `gorm:"type:varchar(255);column:player;" json:"player"`
	CommitTxnHash   string  `gorm:"type:varchar(255);unique;column:commit_transaction;primaryKey" json:"commit_txn_hash"`
	ResponseTxnHash string  `gorm:"type:varchar(255);unique;column:response_transaction;primaryKey" json:"response_txn_hash"`
	Result          bool    `gorm:"type:bool;column:result" json:"result"`
	Value           *BigInt `gorm:"type:varchar(255);column:value;" json:"value"`
	GameId          *BigInt `gorm:"type:varchar(255);column:gameId" json:"gameId"`
	Chain           string  `gorm:"type:varchar(255);column:chain;" json:"chain"`
}

func (*GameResult) TableName() string {
	return "games"
}
