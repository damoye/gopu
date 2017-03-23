package main

import (
	"net/http"

	"golang.org/x/net/websocket"
)

var clients = make(map[string]*websocket.Conn)

func startWebSocket() {
	http.Handle("/websocket", websocket.Handler(handleConn))
}

func handleConn(ws *websocket.Conn) {
	clientKey := ws.Request().Header.Get("id")
	clients[clientKey] = ws
	// TODO
	delete(clients, clientKey)
	ws.Close()
}
