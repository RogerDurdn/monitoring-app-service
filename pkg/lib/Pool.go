package lib

import (
	"fmt"
	"github.com/RogerDurdn/MonitoringApp/pkg/model"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan model.Data
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan model.Data),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Conn Pool: ", len(pool.Clients))
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Conn Pool: ", len(pool.Clients))
			break
		case data := <-pool.Broadcast:
			fmt.Println("Sending data to all clients in Pool size:", len(pool.Clients))
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(data); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
