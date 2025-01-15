package initialize

import (
	"go-ecomerce-backend-api/m/global"
	"go-ecomerce-backend-api/m/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}