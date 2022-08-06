package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// lol, small config
type Config struct {
	RPCHost          string   `json:"host" env:"RPC_HOST"`
	ChainId          int64    `json:"chainId" env:"CHAIN_ID"`
	KeystoreConf     KeyStore `json:"keystore"`
	MoonFlipContract string   `json:"moonflipContract" env:"CONTRACT_ADDRESS"`
	DatabasePath     string   `json:"database" env:"DB_URI"`
}

type KeyStore struct {
	Path     string `env:"URL_PATH"`
	Password string `env:"PASSWORD"`
}

func ReadConfig(r io.Reader) (*Config, error) {
	b, err := ioutil.ReadAll(r)

	if err != nil {
		return nil, err
	}

	var config Config // define an empty config struct

	if err = json.Unmarshal(b, &config); err != nil {
		return nil, err
	}

	return &config, nil // return a nil error and the pointer to our parsed config
}
