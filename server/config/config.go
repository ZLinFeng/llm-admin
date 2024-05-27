package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/BurntSushi/toml"
)

const (
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorReset   = "\033[0m"
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
			fmt.Printf("%sFatal error%s while loading system config file: %s", ColorRed, ColorReset, err)
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
	c.Pattern = strings.ToLower(c.Pattern)
	c.Level = strings.ToLower(c.Level)
	if c.Pattern == "" {
		c.Pattern = "std"
	} else {
		parts := strings.Split(c.Pattern, ",")
		for _, pattern := range parts {
			if pattern != "std" && pattern != "file" {
				fmt.Printf("Only %s`std`%s and %s`file`%s log pattern are supported.",
					ColorRed,
					ColorReset,
					ColorRed,
					ColorReset)
				os.Exit(1)
			}
		}
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
	fmt.Println("Server:")
	fmt.Printf("%20s%16d\n", "Port:", c.Server.Port)
	fmt.Println("Database:")
	fmt.Printf("%20s%16s\n", "Host:", c.Datebase.Host)
	fmt.Printf("%20s%16d\n", "Port:", c.Datebase.Port)
	fmt.Printf("%20s%16s\n", "Username:", c.Datebase.Username)
	fmt.Printf("%20s%16s\n", "Password:", "******")
	fmt.Println("Log:")
	fmt.Printf("%20s%16s\n", "Pattern:", c.Log.Pattern)
	fmt.Printf("%20s%16s\n", "Level:", c.Log.Level)
	fmt.Printf("%20s%16s\n", "File Size:", fmt.Sprintf("%d MB", c.Log.Size))
	fmt.Printf("%20s%16s\n", "Retention Period:", fmt.Sprintf("%d days", c.Log.Days))
}

func GetConfig() *Config {
	if globalConfig == nil {
		log.Fatal("Config not initialized.")
	}
	return globalConfig
}
