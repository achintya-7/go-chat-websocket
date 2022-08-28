package websocket

import (
	"fmt"
)

type Pool struct {
	Register   chan *Client
	UnRegister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client), // to make a new channel
		UnRegister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

// struct method
func (pool *Pool) Start() {
	for {
		select {
			case client := <-pool.Register:
				pool.Clients[client] = true
				fmt.Println("Size of Connection Pool: ", len(pool.Clients))
				for client := range pool.Clients {
					fmt.Println(client)
					client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
				}
				break
			
			case client := <- pool.UnRegister:
				delete(pool.Clients, client)
				fmt.Println("Size of Connection Pool: ", len(pool.Clients))
				for client := range pool.Clients {
					fmt.Println(client)
					client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
				}
				break
			
			case message := <-pool.Broadcast:
				fmt.Println("Sending message to all clients")
				for client := range pool.Clients{
					if err := client.Conn.WriteJSON(message); err != nil {
						fmt.Println(err)
						return
					}
				}
		}
	}
}