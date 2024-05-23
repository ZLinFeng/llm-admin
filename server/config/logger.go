package config

import (
	log "github.com/sirupsen/logrus"
)

func LogSetting(conf *Config) {
	log.SetLevel(log.InfoLevel)
}
