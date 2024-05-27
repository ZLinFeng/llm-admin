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

func printBanner() {
	banner := "    %s__    __    __  ___%s      %s___       __          _%s     \n" +
		"   %s/ /   / /   /  |/  /%s     %s/   | ____/ /___ ___  (_)___%s \n" +
		"  %s/ /   / /   / /|_/ /%s_____%s/ /| |/ __  / __ `__ \\/ / __ \\%s\n" +
		" %s/ /___/ /___/ /  / /%s_____%s/ ___ / /_/ / / / / / / / / / /%s\n" +
		"%s/_____/_____/_/  /_/%s     %s/_/  |_\\__,_/_/ /_/ /_/_/_/ /_/ %s\n"
	fmt.Printf(banner, ColorMagenta, ColorReset, ColorYellow, ColorReset,
		ColorMagenta, ColorReset, ColorYellow, ColorReset,
		ColorMagenta, ColorGreen, ColorYellow, ColorReset,
		ColorMagenta, ColorGreen, ColorYellow, ColorReset,
		ColorMagenta, ColorReset, ColorYellow, ColorReset)
	fmt.Println()
	fmt.Printf(" %s:: %sLLM-Admin%s ::       %s(v%s.RELEASE)%s\n",
		ColorGreen, ColorMagenta, ColorGreen, ColorYellow, "1.0.0", ColorReset)
}

func (c *Config) print() {
	printBanner()
	fmt.Printf("+%s+\n", strings.Repeat("-", 34))
	fmt.Printf("|Server%13s%s%15d%s|\n", "Port:", ColorCyan, c.Server.Port, ColorReset)
	fmt.Printf("+%s+\n", strings.Repeat("-", 34))
	fmt.Printf("|Database%11s%s%15s%s|\n", "Host:", ColorCyan, c.Datebase.Host, ColorReset)
	fmt.Printf("|%19s%s%15d%s|\n", "Port:", ColorCyan, c.Datebase.Port, ColorReset)
	fmt.Printf("|%19s%s%15s%s|\n", "Username:", ColorCyan, c.Datebase.Username, ColorReset)
	fmt.Printf("|%19s%15s|\n", "Password:", "******")
	fmt.Printf("+%s+\n", strings.Repeat("-", 34))
	fmt.Printf("|Log%16s%s%15s%s|\n", "Pattern:", ColorCyan, c.Log.Pattern, ColorReset)
	fmt.Printf("|%19s%s%15s%s|\n", "Level:", ColorCyan, c.Log.Level, ColorReset)
	fmt.Printf("|%19s%s%15s%s|\n", "File Size:", ColorCyan, fmt.Sprintf("%d MB", c.Log.Size), ColorReset)
	fmt.Printf("|%19s%s%15s%s|\n", "Retention Period:", ColorCyan, fmt.Sprintf("%d days", c.Log.Days), ColorReset)
	fmt.Printf("+%s+\n", strings.Repeat("-", 34))
}

func GetConfig() *Config {
	if globalConfig == nil {
		log.Fatal("Config not initialized.")
	}
	return globalConfig
}
