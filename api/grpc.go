package api

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/ptypes/empty"
	"liuqichat/pkg/proto"
	"liuqichat/servers/wb"
	"time"
)

type Chat struct {
	
}

func (c *Chat) SendAllMsg(ctx context.Context, r *proto.SendInfoRequest) (*empty.Empty, error) {
	// 收到群发消息后 发送给 manage 的所有 client
	// 广播消息
	data := map[string]interface{}{
		"type": "sendAll",
		"data": map[string]string{
			"msg": r.Message,
			"uuid": r.Uuid,
			"time": time.Now().String()[6:11]+" "+time.Now().String()[11:19],
		},
	}
	byt, _ := json.Marshal(data)
	wb.GetClientManager().BroadcastMsg(byt)
	return &empty.Empty{}, nil
}