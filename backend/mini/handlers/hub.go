package handlers

// 返回一个hub
func NewHub() *Hub {
	return &Hub{
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

//处理用户连接与断开
func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.register:
			UserRegisterEvent(hub, client)

		case client := <-hub.unregister:
			UserDisconnectEvent(hub, client)
		}
	}
}
