package service

import (
	"errors"

	"github.com/ZlinFeng/llm-admin/server/entity/domain"
	"github.com/ZlinFeng/llm-admin/server/initialize"
	"github.com/ZlinFeng/llm-admin/server/util"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserService struct{}

func (userService *UserService) GetUserByName(username string) (*domain.User, error) {
	var user domain.User
	db := initialize.GetDb()
	err := db.Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Info("User is not exist: " + username)
	}
	return &user, err
}

func (userService *UserService) CheckUserPwd(username string, password string) bool {
	user, err := userService.GetUserByName(username)
	if err == nil {
		return util.CheckPassword(password, user.Password)
	}
	return false
}
