package initialize

import (
	"errors"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/ZlinFeng/llm-admin/server/config"
	gsql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitTidb() *gorm.DB {
	globalC := config.GetConfig()
	globalC.Datebase.Valid()
	dsn := globalC.Datebase.Dsn()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		handleErr(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Error("")
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}

func handleErr(err error) {
	var sqlErr *gsql.MySQLError
	if errors.As(err, &sqlErr) {
		log.Info(sqlErr.Number)
	} else {
		log.Error(err.Error())
	}
}
