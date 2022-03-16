package lib

import (
	"fmt"
	"github.com/RogerDurdn/MonitoringApp/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	wsUpgrade = websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
	}
)

var pool *Pool
var lastData *model.Data

func ChangeData(data model.Data)  {
	fmt.Println(data)
	lastData = &data
	if pool != nil && pool.Broadcast != nil {
		pool.Broadcast <- data
	}
}

func SocketConnection(c *gin.Context)  {
	fmt.Println(c.Request.URL.Query()); clientId := c.Request.URL.Query().Get("clientId")
	fmt.Println("creating a connection for:"+ clientId)
	if pool == nil {
		pool = NewPool()
		pool.Start()
	}
	wsConnection, err := wsUpgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("error connecting to websocket")
	}
	client := NewClient(clientId, wsConnection, pool)
	fmt.Println("Register client")
	pool.Register <- client
	client.Read(lastData)
}
