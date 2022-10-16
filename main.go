package main

import (
	"GooTask/config"
	"GooTask/global"
	"GooTask/logger"
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	logger.InitLogger()
	config.InitServerConfig()
	r := gin.Default()
	// Add a ginzap middleware
	r.Use(ginzap.Ginzap(logger.Logger, time.RFC3339, true))
	// Logs all panic to error log
	r.Use(ginzap.RecoveryWithZap(logger.Logger, true))
	r.Run(fmt.Sprintf(":%v", global.Settings.Port))
}
