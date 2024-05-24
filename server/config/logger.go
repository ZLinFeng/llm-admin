package config

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

type LogFormatter struct{}

func (f *LogFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestampStr := time.Now().Format("2006-01-02 15:04:05")
	var codeFile, funcName string
	var line int
	if entry.Caller != nil {
		codeFile = entry.Caller.File
		line = entry.Caller.Line
		funcName = entry.Caller.Function
		if len(funcName) > 15 {

		}
	}
	return []byte(fmt.Sprintf("%s %s --- [%15s:%4d]%s", timestampStr, entry.Level, codeFile)), nil
}

func LogSetting(conf *Config) {
	log.SetFormatter(&log.TextFormatter{
		DisableColors:          false,
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05",
		DisableLevelTruncation: true,
	})
	log.SetLevel(log.InfoLevel)
	log.Info("Hello World.")
}
