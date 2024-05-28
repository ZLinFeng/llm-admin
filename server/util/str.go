package util

import (
	"fmt"

	c "github.com/ZlinFeng/llm-admin/server/constant"
)

func PrintWithColor(text string, color string) string {
	return fmt.Sprintf("%s%s%s", color, text, c.ColorReset)
}

func PrintWithSuccess(text string) string {
	return PrintWithColor(text, c.ColorGreen)
}

func PrintWithFail(text string) string {
	return PrintWithColor(text, c.ColorRed)
}

func PrintWithWarn(text string) string {
	return PrintWithColor(text, c.ColorYellow)
}
