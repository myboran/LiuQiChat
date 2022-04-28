package wb

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Client struct {
	Addr   string
	Conn   *websocket.Conn
	Send   chan []byte
	UserId string
}

func CreateClient(addr, userId string, conn *websocket.Conn) *Client {
	client := &Client{
		Addr:   addr,
		Conn:   conn,
		Send:   make(chan []byte, 777),
		UserId: userId,
	}
	return client
}

func (c *Client) Read() {
	defer func() {
		println("defer Read")
		c.Conn.Close()
		GetClientManager().Unregister <- c
		close(c.Send)
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			println(fmt.Sprintf("读取客户端数据失败 Addr: %v, err:%v", c.Addr, err))
			return
		}
		// 处理程序
		ProcessingLogin(message, c)
	}
}

func (c *Client) Write() {
	defer func() {
		println("defer Write")
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				println(fmt.Sprintf("读取 c.send 数据失败 Addr: %v", c.Addr))
				return
			}
			err := c.Conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				println(fmt.Sprintf("读取客户端数据失败 Addr: %v, err:%v", c.Addr, err))
			}
		}
	}
}

func (c *Client) SendMsg(msg []byte) {
	c.Send <- msg
}