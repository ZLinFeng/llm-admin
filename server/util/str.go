package util

import (
	"fmt"

	c "github.com/ZlinFeng/llm-admin/server/constant"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Error("Generate password error.")
	}
	return string(bytes)
}

func CheckPassword(password, hashPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(password))
	return err == nil
}
