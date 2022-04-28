package api

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"net/http"

	"github.com/gorilla/websocket"

	"liuqichat/servers/wb"
)

func WsUpgrade(w http.ResponseWriter, r *http.Request) {
	// 升级协议
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}).Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	userId := uuid.NewV4()
	client := wb.CreateClient(conn.RemoteAddr().String(),userId.String() , conn)

	go client.Read()
	go client.Write()
	// 返回uuid
	data := map[string]interface{}{
		"type": "login",
		"data": map[string]interface{}{
			"uuid": userId.String(),
		},
	}
	byt, _ := json.Marshal(data)
	client.SendMsg(byt)
	wb.GetClientManager().RegisterClient(client)
}


