package websocket

//
//import (
//	"log"
//	"net/http"
//	"sync"
//
//	"github.com/gogf/gf/v2/net/ghttp"
//	"github.com/gorilla/websocket"
//)
//
//// 定义 WebSocket 升级器
//var upgrader = websocket.Upgrader{
//	CheckOrigin: func(r *http.Request) bool {
//		return true
//	},
//}
//
//// 存储客户端连接
//var (
//	clients   = make(map[*websocket.Conn]bool)
//	clientsMu sync.Mutex               // 保护 clients map 的互斥锁
//	broadcast = make(chan string, 256) // 广播信息通道，带缓冲
//)
//
//// 初始化广播处理（只启动一次）
//var once sync.Once
//
//// HandlerConnection 处理websocket连接
//func HandlerConnectionV1(r *ghttp.Request) {
//	// 启动广播处理 goroutine（只启动一次）
//	once.Do(func() {
//		go handleBroadcast()
//	})
//
//	ws, err := upgrader.Upgrade(r.Response.Writer, r.Request, nil)
//	if err != nil {
//		log.Printf("upgrade error: %v", err)
//		return
//	}
//	defer func() {
//		clientsMu.Lock()
//		delete(clients, ws)
//		clientsMu.Unlock()
//		ws.Close()
//	}()
//
//	// 注册新客户端
//	clientsMu.Lock()
//	clients[ws] = true
//	clientsMu.Unlock()
//
//	for {
//		var message string
//		err := ws.ReadJSON(&message)
//		if err != nil {
//			log.Printf("read error: %v", err)
//			clientsMu.Lock()
//			delete(clients, ws)
//			clientsMu.Unlock()
//			break
//		}
//		// 将消息发送到广播通道
//		select {
//		case broadcast <- message:
//		default:
//			log.Printf("broadcast channel full, dropping message")
//		}
//	}
//}
//
//// 将广播通道给所有的用户
//func handleBroadcast() {
//	for {
//		message := <-broadcast
//		clientsMu.Lock()
//		for client := range clients {
//			err := client.WriteJSON(message)
//			if err != nil {
//				log.Printf("write error: %v", err)
//				client.Close()
//				delete(clients, client)
//			}
//		}
//		clientsMu.Unlock()
//	}
//}
