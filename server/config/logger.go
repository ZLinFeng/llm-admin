package config

import (
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type LogFormatter struct {
	pattern string
}

const (
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
	ColorReset   = "\033[0m"
)

func (f *LogFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestampStr := time.Now().Format("2006-01-02 15:04:05")
	var codeFile, funcName, color string
	var line int
	if entry.Caller != nil {
		codeFile = entry.Caller.File
		line = entry.Caller.Line
		funcName = entry.Caller.Function

		parts := strings.Split(funcName, ".")
		funcName = parts[len(parts)-1]
		// ASCII 字节|字符长度
		if len(funcName) > 15 {
			slice := []rune(funcName)
			slice = slice[0:12]
			funcName = string(slice) + "..."
		}

		if len(codeFile) > 30 {
			slice := []rune(codeFile)
			slice = slice[len(codeFile)-27 : len(codeFile)]
			codeFile = "..." + string(slice)
		}
		switch entry.Level {
		case log.WarnLevel:
			color = ColorYellow
		case log.ErrorLevel:
			color = ColorRed
		case log.FatalLevel:
			color = ColorRed
		default:
			color = ColorGreen
		}
	}
	return []byte(fmt.Sprintf("%s%s%s %s%s%s --- %s[%15s:%4d]%s %s%30s:%s %s\n",
		ColorCyan,
		timestampStr,
		ColorReset,
		color,
		entry.Level,
		ColorReset,
		ColorMagenta,
		funcName,
		line,
		ColorReset,
		ColorBlue,
		codeFile,
		ColorReset,
		entry.Message)), nil
}

func LogSetting(conf *Config) {
	log.SetFormatter(new(LogFormatter))
	log.SetReportCaller(true)
	log.SetLevel(log.InfoLevel)
	log.Info("Hello World.")
}
