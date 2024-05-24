package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server   ServerConfig   `toml:"server"`
	Datebase DatabaseConfig `toml:"database"`
	Log      LogConfig      `toml:"log"`
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

type LogConfig struct {
	Pattern string `toml:"pattern"`
	Level   string `toml:"level"`
	Size    int    `toml:"size"`
	Days    int    `toml:"days"`
}

var globalConfig *Config
var once sync.Once

func LoadSysSetting() {
	once.Do(func() {
		globalConfig = &Config{}
		_, err := toml.DecodeFile("config.toml", globalConfig)
		if err != nil {
			fmt.Printf("Fatal error while loading system config file: %s", err)
			os.Exit(1)
		}
		// 校验并设置默认字段
		globalConfig.Server.valid()
		globalConfig.Datebase.valid()
		globalConfig.Log.valid()
		globalConfig.print()
	})
}

func (c *ServerConfig) valid() {
	if c.Port == 0 {
		c.Port = 8080
	}
}

func (c *DatabaseConfig) valid() {
	if c.Host == "" {
		c.Host = "127.0.0.1"
	}
	if c.Port == 0 {
		c.Port = 4000
	}
	if c.Username == "" {
		c.Username = "root"
	}
}

func (c *LogConfig) valid() {
	if c.Pattern == "" {
		c.Pattern = "std"
	}
	if c.Days == 0 {
		c.Days = 7
	}
	if c.Level == "" {
		c.Level = "info"
	}
	if c.Size == 0 {
		c.Size = 100
	}
}

func (c *Config) print() {
	fmt.Println()
}

func GetConfig() *Config {
	if globalConfig == nil {
		log.Fatal("Config not initialized.")
	}
	return globalConfig
}
