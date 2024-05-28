package config

import "fmt"

type DatabaseConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Dbname   string `toml:"Dbname"`
}

func (t *DatabaseConfig) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf-8mb4&parseTime=True&loc=Local",
		t.Username, t.Password, t.Host, t.Port, t.Dbname)
}

func (t *DatabaseConfig) Valid() {

}
