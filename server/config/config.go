package config

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/BurntSushi/toml"
	c "github.com/ZlinFeng/llm-admin/server/constant"
)

type Config struct {
	Server   ServerConfig   `toml:"server"`
	Datebase DatabaseConfig `toml:"database"`
	Log      LogConfig      `toml:"log"`
}

type ServerConfig struct {
	Port int `toml:"port"`
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
			fmt.Printf("%sFatal error%s while loading system config file: %s",
				c.ColorRed,
				c.ColorReset, err)
			os.Exit(1)
		}
		// 设置默认字段
		globalConfig.Server.defaultValue()
		globalConfig.Datebase.defaultValue()
		globalConfig.Log.defaultValue()
		globalConfig.print()
	})
}

func (c *ServerConfig) defaultValue() {
	if c.Port == 0 {
		c.Port = 8080
	}
}

func (c *DatabaseConfig) defaultValue() {
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

func (f *LogConfig) defaultValue() {
	f.Pattern = strings.ToLower(f.Pattern)
	f.Level = strings.ToLower(f.Level)
	if f.Pattern == "" {
		f.Pattern = "std"
	} else {
		parts := strings.Split(f.Pattern, ",")
		for _, pattern := range parts {
			if pattern != "std" && pattern != "file" {
				fmt.Printf("Only %s`std`%s and %s`file`%s log pattern are supported.",
					c.ColorRed,
					c.ColorReset,
					c.ColorRed,
					c.ColorReset)
				os.Exit(1)
			}
		}
	}
	if f.Days == 0 {
		f.Days = 7
	}
	if f.Level == "" {
		f.Level = "info"
	}
	if f.Size == 0 {
		f.Size = 100
	}
}

func printBanner() {
	banner := "    %s__    __    __  ___%s      %s___       __          _%s     \n" +
		"   %s/ /   / /   /  |/  /%s     %s/   | ____/ /___ ___  (_)___%s \n" +
		"  %s/ /   / /   / /|_/ /%s_____%s/ /| |/ __  / __ `__ \\/ / __ \\%s\n" +
		" %s/ /___/ /___/ /  / /%s_____%s/ ___ / /_/ / / / / / / / / / /%s\n" +
		"%s/_____/_____/_/  /_/%s     %s/_/  |_\\__,_/_/ /_/ /_/_/_/ /_/ %s\n"
	fmt.Printf(banner, c.ColorMagenta, c.ColorReset, c.ColorYellow, c.ColorReset,
		c.ColorMagenta, c.ColorReset, c.ColorYellow, c.ColorReset,
		c.ColorMagenta, c.ColorGreen, c.ColorYellow, c.ColorReset,
		c.ColorMagenta, c.ColorGreen, c.ColorYellow, c.ColorReset,
		c.ColorMagenta, c.ColorReset, c.ColorYellow, c.ColorReset)
	fmt.Println()
	fmt.Printf(" %s:: %sLLM-Admin%s ::       %s(v%s.RELEASE)%s\n",
		c.ColorGreen, c.ColorMagenta, c.ColorGreen, c.ColorYellow, "1.0.0", c.ColorReset)
}

func (f *Config) print() {
	printBanner()
	fmt.Printf("+%s+\n", strings.Repeat("-", 34))
	fmt.Printf("|Server%13s%s%15d%s|\n", "Port:", c.ColorCyan, f.Server.Port, c.ColorReset)
	fmt.Printf("+%s+\n", strings.Repeat("-", 34))
	fmt.Printf("|Database%11s%s%15s%s|\n", "Host:", c.ColorCyan, f.Datebase.Host, c.ColorReset)
	fmt.Printf("|%19s%s%15d%s|\n", "Port:", c.ColorCyan, f.Datebase.Port, c.ColorReset)
	fmt.Printf("|%19s%s%15s%s|\n", "Username:", c.ColorCyan, f.Datebase.Username, c.ColorReset)
	fmt.Printf("|%19s%15s|\n", "Password:", "******")
	fmt.Printf("+%s+\n", strings.Repeat("-", 34))
	fmt.Printf("|Log%16s%s%15s%s|\n", "Pattern:", c.ColorCyan, f.Log.Pattern, c.ColorReset)
	fmt.Printf("|%19s%s%15s%s|\n", "Level:", c.ColorCyan, f.Log.Level, c.ColorReset)
	fmt.Printf("|%19s%s%15s%s|\n", "File Size:", c.ColorCyan, fmt.Sprintf("%d MB", f.Log.Size), c.ColorReset)
	fmt.Printf("|%19s%s%15s%s|\n", "Retention Period:", c.ColorCyan, fmt.Sprintf("%d days", f.Log.Days), c.ColorReset)
	fmt.Printf("+%s+\n", strings.Repeat("-", 34))
}

func GetConfig() *Config {
	if globalConfig == nil {
		log.Fatal("Config not initialized.")
	}
	return globalConfig
}
