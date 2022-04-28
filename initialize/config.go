package initialize

import (
	"github.com/spf13/viper"
	"liuqichat/global"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile("./config/config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.UnmarshalKey("Server", &global.ServerInfo); err != nil {
		panic(err)
	}
	if err := v.UnmarshalKey("WebSocket", &global.WebSocketInfo); err != nil {
		panic(err)
	}
	if err := v.UnmarshalKey("Mysql", &global.MysqlInfo); err != nil {
		panic(err)
	}
	if err := v.UnmarshalKey("Redis", &global.RedisInfo); err != nil {
		panic(err)
	}
	if err := v.UnmarshalKey("GRPC", &global.GRPCInfo); err != nil {
		panic(err)
	}
}
