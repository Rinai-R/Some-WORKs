package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]string)
var mutex = &sync.Mutex{}

// 处理连接
func handleConnection(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	// 生成唯一ID
	clientID := uuid.New().String()

	mutex.Lock()
	clients[conn] = clientID // 将连接和ID保存
	mutex.Unlock()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		log.Println("收到消息:", string(msg))
		broadcastMessage(msg, conn) // 传递当前连接以排除发送者
	}

	mutex.Lock()
	delete(clients, conn) // 移除此连接
	mutex.Unlock()
}

// 广播消息
func broadcastMessage(msg []byte, sender *websocket.Conn) {
	mutex.Lock()
	defer mutex.Unlock()

	for client := range clients {
		if client != sender { // 跳过发送者
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("Write error:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	r := gin.Default()
	r.StaticFile("/", "./index.html")
	r.GET("/chat", handleConnection)
	r.Run(":8080")
}
