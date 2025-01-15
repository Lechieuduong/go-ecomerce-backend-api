package initialize

import (
	"fmt"
	"go-ecomerce-backend-api/m/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	m := global.Config
	fmt.Println("Loading configuration mysql: ", m.Mysql.Username, m.Mysql.Password)
	InitLogger()
	global.Logger.Info("Config Log ok!", zap.String("ok", "success"))
	InitMysql()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}