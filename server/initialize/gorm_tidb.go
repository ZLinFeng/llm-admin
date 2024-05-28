package initialize

import (
	"github.com/ZlinFeng/llm-admin/server/config"
	"gorm.io/gorm"
)

func InitTidb() *gorm.DB {
	globalC := config.GetConfig()

	dsn := globalC.Datebase.Dsn()
}
