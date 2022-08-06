package main

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/big"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/MoonFlipDevelopment/moonflip/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/storage/redis"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	conf  *config.Config
	store *keystore.KeyStore

	db *gorm.DB

	txnPattern = regexp.MustCompile("(0x)?([a-fA-F0-9]{64})")

	txnRoutePattern = regexp.MustCompile(`(\/process_txn\/)(0x)?([a-fA-F0-9]{64})`)

	privateKey *keystore.Key

	createKeyFlag = flag.Bool("createKey", false, "creates a instead of running the application")
)

func CreateKey(password string) (string, error) {

	ks := keystore.NewKeyStore("keys", keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := ks.NewAccount(password)

	if err != nil {
		return "", err
	}

	jsonKey, err := ks.Export(account, password, password)

	if err != nil {
		return "", err
	}

	return string(jsonKey), err
}

const (
	Heads = 0
	Tails = 1

	DatabaseTable = "transactions"
	GameTable     = "games"
)

func uriToMySQLConfig(uri string) (*mysqlDriver.Config, error) {
	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	mysqlConfig := mysqlDriver.NewConfig()

	mysqlConfig.Addr = u.Host
	mysqlConfig.DBName = u.Path[1:]
	mysqlConfig.Net = "tcp"
	mysqlConfig.User = u.User.Username()

	mysqlConfig.ParseTime = true

	passwd, present := u.User.Password()

	if present {
		mysqlConfig.Passwd = passwd
	}

	return mysqlConfig, nil
}

func initKey() *keystore.Key {
	decoded, err := base64.StdEncoding.DecodeString(os.Getenv(PRIVATE_KEY))

	if err != nil {
		log.Fatalln("failed to decode base64 string:", err)
	}

	key, err := keystore.DecryptKey(decoded, os.Getenv(PASSWORD))

	if err != nil {
		log.Fatalln("failed to decrypt key:", err)
	}

	log.Printf("using account %s, make sure this is well-funded!\n", key.Address.String())
	return key
}

func initDatabase() *gorm.DB {

	dbEngine := os.Getenv(DB_ENGINE)

	if dbEngine == "" {
		dbEngine = "sqlite"
	}

	log.Printf("Using Database Engine %s", dbEngine)

	var db *gorm.DB
	var err error

	if dbEngine == "sqlite" {
		db, err = gorm.Open(sqlite.Open(os.Getenv(DB_URI)), &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		})
	} else if dbEngine == "mysql" {
		mysqlDriverConfig, _ := uriToMySQLConfig(os.Getenv(DB_URI))

		db, err = gorm.Open(mysql.Open(mysqlDriverConfig.FormatDSN()), &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		})

	}
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func init() {
	flag.Parse()

	if *createKeyFlag {
		return
	}
	log.Println("loading private key...")
	privateKey = initKey()

	log.Println("initializing database...")
	db = initDatabase()

	log.Println("running database migrations...")
	db.AutoMigrate(&Transaction{})
	db.AutoMigrate(&GameResult{})
}

func SignerFunction(chainId *big.Int) bind.SignerFn {
	return func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		signer := types.LatestSignerForChainID(chainId)
		return types.SignTx(tx, signer, privateKey.PrivateKey)
	}
}

func getHex(txnString string) string {
	if txnPattern.MatchString(txnString) {
		transaction := txnPattern.FindAllStringSubmatch(txnString, -1)
		return transaction[0][2]
	}
	return ""
}

func main() {

	if *createKeyFlag {
		newAccount, err := CreateKey(os.Getenv(PASSWORD))

		if err != nil {
			log.Fatalln(err)
		}

		log.Println(newAccount)
		return
	}

	// this too has been moved elsewhere. TODO: Fix hardcoded strings, include envvars like %s_RPC_HOST where %s = chain symbol
	// rpcClient, err := ethclient.Dial(os.Getenv(RPC_HOST))

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	//
	redisStorage := redis.New(redis.Config{
		URL:       os.Getenv(REDIS_URI),
		TLSConfig: &tls.Config{},
	})

	// Redis will be used for discord bot but not atm
	// options, _ := redisDriver.ParseURL(os.Getenv(REDIS_URI))
	// options.TLSConfig = &tls.Config{}
	// redisClient := redisDriver.NewClient(options)

	// this has been moved elsewhere.
	// moonflipContract := common.HexToAddress(os.Getenv(CONTRACT_ADDRESS))

	// moonflipProgram, err := moonflip.NewMoonflip(moonflipContract, rpcClient)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	primaryAccount := privateKey

	app := fiber.New()

	app.Use(logger.New())

	app.Use(limiter.New(limiter.Config{
		Max: 1,
		Next: func(c *fiber.Ctx) bool {
			path := c.Path()

			if txnRoutePattern.MatchString(path) {

				matches := txnRoutePattern.FindAllStringSubmatch(path, -1)
				txn := fmt.Sprintf("%s%s", matches[0][2], matches[0][3])

				// this is used for DB lookups.
				if !strings.HasPrefix(txn, "0x") {
					txn = "0x" + txn
				}
				// check if we've already received a transcation
				var exists bool

				ExistsInDatabase(DatabaseTable, db, db.Table(DatabaseTable).Where("transaction_hash = ?", txn)).Scan(&exists)

				return exists
			}

			return true
		},
		Expiration: 30 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			path := c.Path()

			if txnRoutePattern.MatchString(path) {
				matches := txnRoutePattern.FindAllStringSubmatch(path, -1)

				key := fmt.Sprintf("%s%s", matches[0][2], matches[0][3])

				if !strings.HasPrefix(key, "0x") {
					key = "0x" + key
				}

				return key + ":" + c.IP()
			}

			return c.IP()
		},
		Storage: redisStorage,
	}))

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: os.Getenv(ALLOWED_ORIGIN),
	// 	AllowMethods: "POST",
	// }))

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	app.Post("/process_txn/movr/:transaction", GetChain("movr").HandleTransaction(primaryAccount))
	app.Post("/process_txn/avax/:transaction", GetChain("avax").HandleTransaction(primaryAccount))

	if os.Getenv("DEVNETS") != "" {
		app.Post("/process_txn/fuji/:transaction", GetChain("fuji").HandleTransaction(primaryAccount))
	}

	log.Println(app.Listen(os.Getenv(WEB_HOST)))
}

func Ecrecover(messageHash, signature []byte) (*common.Address, error) {
	if len(signature) != crypto.SignatureLength {
		return nil, errors.New("invalid signature")
	}

	recoveryId := signature[crypto.RecoveryIDOffset]

	log.Println(recoveryId)

	if recoveryId != 27 && recoveryId != 28 {
		return nil, errors.New("unable to recover signature!")
	}

	signature[crypto.RecoveryIDOffset] -= 27

	publicKey, err := crypto.SigToPub(messageHash, signature)

	if err != nil {
		return nil, err
	}

	recoveredAddress := crypto.PubkeyToAddress(*publicKey)
	return &recoveredAddress, nil
}
