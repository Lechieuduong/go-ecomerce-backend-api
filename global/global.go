package global

import (
	"go-ecomerce-backend-api/m/pkg/logger"
	"go-ecomerce-backend-api/m/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb *gorm.DB
)