package initialize

import (
	"fmt"
	"go-ecomerce-backend-api/m/global"
)

func Run() {
	LoadConfig()
	m := global.Config
	fmt.Println("Loading configuration mysql: ", m.Mysql.Username, m.Mysql.Password)
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}