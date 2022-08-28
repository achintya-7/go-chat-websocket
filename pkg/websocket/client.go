package websocket

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID      string
	Conn    *websocket.Conn
	Pool    *Pool
	mu      sync.Mutex
}

type Message struct {
	Type    int      `json:"type"`
	Body    string   `json:"body"`
}

// struct method
func (c *Client) Read() {
	
	// * defer function will run at the end of function call
	defer func () {
		c.Pool.UnRegister <- c
		c.Conn.Close()
	}() // anonymous function

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Printf("Message Recieved\n Type : %+v\n Body : %+v", message.Type, message.Body)
	}
	
}