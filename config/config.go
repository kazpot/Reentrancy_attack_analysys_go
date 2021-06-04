package config

import (
	"gopkg.in/ini.v1"
	"log"
)

type ConfigList struct {
	Key             string
	Password        string
	Infura          string
	NetworkId       int64
	LogFile         string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read config file: %v\n", err)
	}

	Config = ConfigList{
		Key:                cfg.Section("rinkeby").Key("key").String(),
		Password:           cfg.Section("rinkeby").Key("password").String(),
		Infura:             cfg.Section("rinkeby").Key("infura").String(),
		NetworkId:          cfg.Section("rinkeby").Key("network_id").MustInt64(),
		LogFile:            cfg.Section("rinkeby").Key("log_file").String(),
	}

}