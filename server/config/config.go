package config

import (
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
			log.Fatalf("Fatal error while loading system config file: %s", err)
			os.Exit(1)
		}
		// 校验并设置默认字段
		globalConfig.Server.valid()
		globalConfig.Datebase.valid()
		globalConfig.Log.valid()
	})
}

func (c *ServerConfig) valid() {
	if c.Port == 0 {
		log.Printf("Use default server port: 8080")
		c.Port = 8080
	}
}

func (c *DatabaseConfig) valid() {

}

func (c *LogConfig) valid() {

}

func GetConfig() *Config {
	if globalConfig == nil {
		log.Fatal("Config not initialized.")
	}
	return globalConfig
}
