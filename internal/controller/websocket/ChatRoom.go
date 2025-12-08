package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gorilla/websocket"
)

// ChatMessage 聊天消息结构
type ChatMessage struct {
	RoomId  string `json:"roomId"`
	UserId  string `json:"userId"`
	Message string `json:"message"`
	Type    string `json:"type"` // join, leave, message
}

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan ChatMessage, 256) // 使用带缓冲的通道
var mutex = sync.RWMutex{}
var once sync.Once // 确保广播 goroutine 只启动一次

// HandleChatConnections 处理 WebSocket 连接（适配 GoFrame）
func HandleChatConnections(r *ghttp.Request) {
	// 启动广播处理 goroutine（只启动一次）
	once.Do(func() {
		go handleBroadcast()
	})

	ws, err := upgrader.Upgrade(r.Response.Writer, r.Request, nil)
	if err != nil {
		log.Printf("upgrade error: %v", err)
		return
	}

	mutex.Lock()
	clients[ws] = true
	mutex.Unlock()

	defer func() {
		mutex.Lock()
		delete(clients, ws)
		mutex.Unlock()
		ws.Close()
	}()

	for {
		var msg ChatMessage
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("read error: %v", err)
			break
		}
		// 将消息发送到广播通道
		select {
		case broadcast <- msg:
		default:
			log.Printf("broadcast channel full, dropping message")
		}
	}
}

// handleBroadcast 处理消息广播
func handleBroadcast() {
	for {
		msg := <-broadcast
		mutex.RLock()
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("write error: %v", err)
				client.Close()
				mutex.RUnlock()
				mutex.Lock()
				delete(clients, client)
				mutex.Unlock()
				mutex.RLock()
			}
		}
		mutex.RUnlock()
	}
}
