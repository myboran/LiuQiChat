package wb

import "sync"

var (
	clientManager *ClientManager
)

type ClientManager struct {
	Clients     map[string]*Client
	ClientsLock sync.RWMutex
	Register    chan *Client
	Unregister  chan *Client
	Broadcast   chan []byte
}

func GetClientManager() *ClientManager {
	if clientManager != nil {
		return clientManager
	}
	clientManager = &ClientManager{
		Clients: make(map[string]*Client),
		Register: make(chan *Client, 777),
		Unregister: make(chan *Client, 777),
		Broadcast: make(chan []byte, 777),
	}
	return clientManager
}

func (m *ClientManager) xxx() {

}

func (m *ClientManager) RegisterClient(client *Client) {
	m.Register <- client
}

func (m *ClientManager) BroadcastMsg(msg []byte) {
	m.Broadcast <- msg
}

func (m *ClientManager) Start() {
	for {
		select {
		case client := <- m.Register:
			m.EventRegister(client)
		case client := <- m.Unregister:
			m.EventUnregister(client)
		case msg := <- m.Broadcast:
			// 广播
			for _, v := range m.Clients {
				v.SendMsg(msg)
			}
		}
	}
}

func (m *ClientManager) EventRegister(client *Client) {
	m.ClientsLock.Lock()
	defer m.ClientsLock.Unlock()

	m.Clients[client.UserId] = client
}

func (m *ClientManager) EventUnregister(client *Client) {
	m.ClientsLock.Lock()
	defer m.ClientsLock.Unlock()

	delete(m.Clients, client.UserId)
}