package initialize

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/ZlinFeng/llm-admin/server/config"
	"github.com/ZlinFeng/llm-admin/server/entity/domain"
	"github.com/ZlinFeng/llm-admin/server/util"
	log "github.com/sirupsen/logrus"

	gsql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var globalDb *gorm.DB
var once sync.Once

func InitTidb() {
	once.Do(func() {
		globalC := config.GetConfig()
		globalC.Datebase.Valid()
		dsn := globalC.Datebase.Dsn()

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			handleErr(err)
		}
		sqlDB, err := db.DB()
		if err != nil {
			log.Error("Fatal error while getting db: " + err.Error())
			os.Exit(1)
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
		log.Info("Connectting db " + util.PrintWithSuccess("[Successed]"))
		globalDb = db
		autoMigrate()
	})
}

func autoMigrate() {
	err := globalDb.AutoMigrate(
		domain.User{},
		domain.Role{},
	)
	if err != nil {
		log.Error("Failed to register tables: " + err.Error())
		os.Exit(0)
	}
	log.Info("Register tables " + util.PrintWithSuccess("[Successed]"))
}

func handleErr(err error) {
	var sqlErr *gsql.MySQLError
	if errors.As(err, &sqlErr) {
		log.Error(fmt.Sprintf("DB Error Code: %d, Error msg: %s", sqlErr.Number, sqlErr.Message))
	} else {
		log.Error(err.Error())
	}
	os.Exit(1)
}

func GetDb() *gorm.DB {
	if globalDb == nil {
		log.Fatal("Database not initialized.")
	}
	return globalDb
}
