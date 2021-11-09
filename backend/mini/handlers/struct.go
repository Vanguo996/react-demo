package handlers

import "github.com/gorilla/websocket"

type Hub struct {
	register   chan *Client
	unregister chan *Client
	clients    map[*Client]bool
	// Broadcast  chan Message
}

type Client struct {
	hub                 *Hub
	webSocketConnection *websocket.Conn
	message             chan SocketEventStruct
	username            string
}

type SocketEventStruct struct {
	EventName    string
	EventPayload string
}
