package main

import (
	"fmt"
	"net/http"

	"github.com/achintya-7/go-chat-websocket/pkg/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request)  {
	fmt.Println("websocket endpoint reached")

	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read() // Read is a struct method

}

func setUpRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})



}

func main() {
	fmt.Println("Chat Socket Started at - ws://localhost:2205/ws")
	setUpRoutes()
	http.ListenAndServe(":2205", nil)
}