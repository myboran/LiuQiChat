package main

import (
	"liuqichat/global"
	"liuqichat/initialize"
)

func main() {
	initialize.InitConfig()
	initialize.InitMysql()
	initialize.InitRedis()
	// 启动websocket服务
	go initialize.InitWebSocket()
	// 启动grpc服务
	go initialize.InitGRPC()

	r := initialize.InitRouter()
	err := r.Run("0.0.0.0:" + global.ServerInfo.Port)
	if err != nil {
		panic(err)
	}

}
