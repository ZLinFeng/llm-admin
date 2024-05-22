package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server   ServerConfig   `toml:"server"`
	Datebase DatabaseConfig `toml:"database"`
}

type ServerConfig struct {
	Port int `toml:"port"`
}

type DatabaseConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

func LoadSysSetting() *Config {
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Fatalf("Fatal error while load system config file: %s", err)
	}

	return &config
}
