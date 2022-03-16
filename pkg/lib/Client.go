package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/RogerDurdn/MonitoringApp/pkg/model"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

func (c *Client) Read(data *model.Data) {

	fmt.Println("client on reading")
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	if data != nil {
		fmt.Println("writing data:", data)
		dataBytes := new(bytes.Buffer)
		err := json.NewEncoder(dataBytes).Encode(data)
		if err != nil {
			fmt.Println("cannot send data on subscription:", c.ID)
			return
		}
		_ = c.Conn.WriteJSON(data)
	}

	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		//var mess Data
		//_ = json.Unmarshal(p, &mess)
		//message := Data{ClientId: c.ID, Message:mess.Message}
		//c.Pool.Broadcast <- message
		fmt.Printf("Data Received: %+v\n", p)
	}
}

func NewClient(id string, c *websocket.Conn, p *Pool) *Client {
	newClient :=  Client{
		ID:   id,
		Conn: c,
		Pool: p,
	}
	return &newClient
}


