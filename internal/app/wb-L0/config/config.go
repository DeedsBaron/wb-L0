package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"wb-L0/internal/app/wb-L0/flags"
)

var (
	Config config
)

func init() {
	err := NewConfig(&Config)
	if err != nil {
		log.Fatal("Can't parse config: ", err.Error())
	}
}

type config struct {
	BindAddr           string `toml:"bind_addr"`
	LogLevel           string `toml:"log_level"`
	Token              string `toml:"token"`
	RequestFreq        int    `toml:"requestfreq"`
	SendMessageTimeout int    `toml:"sendmessagetimeout"`

	Storage struct {
		Host     string `toml:"host"`
		Port     string `toml:"port"`
		Database string `toml:"database"`
		Username string `toml:"username"`
		Password string `toml:"password"`
		Attempts int    `toml:"attempts2con"`
	} `toml:"storage"`

	Nats struct {
		ServerID string `toml:"serverID"`
		ClientID string `toml:"clientID"`
		NatsUrl  string `toml:"natsUrl"`
	}
}

func NewConfig(cfg *config) error {
	_, err := toml.DecodeFile(flags.ConfigPath, cfg)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
