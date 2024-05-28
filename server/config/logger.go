package config

import (
	"fmt"
	"strings"
	"time"

	c "github.com/ZlinFeng/llm-admin/server/constant"
	"github.com/ZlinFeng/llm-admin/server/util"
	log "github.com/sirupsen/logrus"
)

type LogFormatter struct {
	pattern string
}

func (f *LogFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestampStr := time.Now().Format("2006-01-02 15:04:05")
	var codeFile, funcName, color string
	var line int
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
		color = c.ColorYellow
	case log.ErrorLevel:
		color = c.ColorRed
	case log.FatalLevel:
		color = c.ColorRed
	default:
		color = c.ColorGreen
	}
	levelWithColor := util.PrintWithColor(strings.ToUpper(entry.Level.String()), color)
	funcAndLine := util.PrintWithColor(fmt.Sprintf("[%15s:%4d]", funcName, line),
		c.ColorMagenta)
	return []byte(fmt.Sprintf("%s %s %s %s %s: %s\n",
		util.PrintWithColor(timestampStr, c.ColorCyan),
		levelWithColor,
		strings.Repeat("-", 8-len(entry.Level.String())),
		funcAndLine,
		util.PrintWithColor(codeFile, c.ColorBlue),
		entry.Message)), nil
}

func LogSetting(conf *Config) {
	formatter := LogFormatter{
		pattern: conf.Log.Pattern,
	}
	log.SetFormatter(&formatter)
	log.SetReportCaller(true)
	log.SetLevel(log.InfoLevel)
	log.Info("Hello World.")
}
