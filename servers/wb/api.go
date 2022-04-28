package wb

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"liuqichat/global"
	"liuqichat/pkg/proto"
	"log"
	"time"
)


type WsReq struct {
	Uuid string `json:"uuid"`
	Action string `json:"action"`
	Msg  string `json:"msg"`
}

func ProcessingLogin(msg []byte, conn *Client)  {
	req := &WsReq{}
	err := json.Unmarshal(msg, req)
	if err != nil {
		conn.SendMsg([]byte("数据不合法"))
		return
	}
	action := req.Action
	switch action {
	case "sendAll":
		sendAll(req, conn)
	case "getChatLog":
		getChatLog(conn)
	default:
		conn.SendMsg([]byte("数据不合法"))
	}
}

func sendAll(req *WsReq, conn *Client) {
	// 广播消息
	data := map[string]interface{}{
		"type": "sendAll",
		"data": map[string]string{
			"msg": req.Msg,
			"uuid": conn.UserId,
			"time": time.Now().String()[6:11]+" "+time.Now().String()[11:19],
		},
	}
	byt, _ := json.Marshal(data)
	go global.SaveSendAllLog(req.Uuid, req.Msg)
	// 通知其他服务器发送消息
	val, err := global.RedisDB.Do(context.Background(), "smembers", "grpcIp").Result()
	if err != nil {
		if err == redis.Nil {
			return
		}
		panic(err)
	}
	for _, v := range val.([]interface{}) {
		// 自己pass

		if fmt.Sprintf("%v", v) == global.GRPCInfo.Host+":"+global.GRPCInfo.Port {
			continue
		}
		go sendMsgToOthers(v, conn.UserId, req.Msg)
	}

	GetClientManager().BroadcastMsg(byt)
}

type ChatLog struct {
	UUID string `json:"uuid"`
	Msg string `json:"msg"`
	Time string `json:"time"`
}

type ChatLogs struct {
	Type string `json:"type"`
	Data []ChatLog `json:"data"`
}

func getChatLog(conn *Client) {
	// 查找聊天记录
	times := time.Now()
	times = times.Add(-7*time.Hour*24)
	SendAllLogs := []global.SendAllLog{}
	global.MysqlConn.Where("time >= ?", times).Order("time desc").Find(&SendAllLogs)

	datas := ChatLogs{Type: "getChatLog"}
	for _, v := range SendAllLogs {
		data := ChatLog{UUID: v.Uuid, Msg: v.Msg, Time: v.Time.String()[6:11]+" "+v.Time.String()[11:19]}
		datas.Data = append(datas.Data, data)
	}
	byt, _ := json.Marshal(datas)
	conn.SendMsg(byt)
}

func sendMsgToOthers(ip interface{}, uuid, msg string) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%v", ip), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("sendMsgToOthers grpc.Dial err: %s", err)
	}
	defer conn.Close()
	c := proto.NewChatClient(conn)
	_, err = c.SendAllMsg(context.Background(), &proto.SendInfoRequest{
		Uuid: uuid,
		Message: msg,
	})
	if err != nil {
		log.Fatalf("sendMsgToOthers SendAllMsg err: %s", err)
	}
}