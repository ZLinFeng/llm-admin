package main

import (
	"log"

	"github.com/ZlinFeng/llm-admin/server/config"
	"github.com/gin-gonic/gin"
)

func init() {
	// 加载系统配置
	config.LoadSysSetting()
	// 配置日志
	config.LogSetting(config.GetConfig())
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		log.Printf("Get request from: %s", c.RemoteIP())
		c.JSON(200, gin.H{
			"res": "pong",
		})
	})
	r.Run()
}
