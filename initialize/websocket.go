package initialize

import (
	"fmt"
	"liuqichat/api"
	"liuqichat/global"
	"liuqichat/servers/wb"
	"net/http"
)

func InitWebSocket() {
	clientManager := wb.GetClientManager()
	http.HandleFunc("/acc", api.WsUpgrade)
	go clientManager.Start()
	fmt.Println("webSocket 启动成功")
	err := http.ListenAndServe("0.0.0.0:"+global.WebSocketInfo.Port, nil)
	if err != nil {
		panic(fmt.Errorf("启动 websocket 服务失败 err: %v", err))
	}
}