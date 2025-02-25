package main

import (
	CustomWebSocket "chatapplication/websocket"
	"log"
	"net/http"
)

var pool *CustomWebSocket.Pool

func serverWs(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := CustomWebSocket.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer conn.Close()

	client := &CustomWebSocket.Client{
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool = CustomWebSocket.NewPool()
	go pool.Start()
	http.HandleFunc("/ws", serverWs)
}

func main() {
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}
