package config

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"wb-L0/internal/app/wb-L0/flags"
	"wb-L0/internal/app/wb-L0/logger"
)

var (
	Config Cfg
)

func init() {
	err := NewConfig(&Config)
	if err != nil {
		logrus.Fatal("Can't parse config: ", err.Error())
	}
	logrus.Info("Config successfully load")
	logger.ConfigureLogger(Config.LogLevel)
}

type Cfg struct {
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
		Interval int    `toml:"interval"`
		MaxOut   int    `toml:"maxout"`
	}
}

func NewConfig(cfg *Cfg) error {
	_, err := toml.DecodeFile(flags.ConfigPath, cfg)
	if err != nil {
		logrus.Fatal(err)
		return err
	}
	return nil
}
