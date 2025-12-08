package websocket

//
//import (
//	"github.com/gorilla/websocket"
//	"log"
//	"net/http"
//	"sync"
//)
//
//var upgraderv2 = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
//var clientsV2 = make(map[*websocket.Conn]bool) //
//var mutex = sync.RWMutex{}                     // 读写锁 // 保护 clientsV2 并发访问
//
////var bufferPool = sync.Pool{
////	New: func() interface{} {
////		return make([]byte, 1024)
////	},
////}
//
//var bufferPool = sync.Pool{
//	New: func() interface{} {
//		return make([]byte, 1024)
//	},
//}
//
//func readMessage(ws *websocket.Conn) ([]byte, error) {
//	buf := bufferPool.Get().([]byte)
//	defer bufferPool.Put(buf)
//	_, data, err := ws.ReadMessage()
//	return data, err
//}
//
//func HandlerConnectionV2(w http.ResponseWriter, r *http.Request) {
//	ws, err := upgraderv2.Upgrade(w, r, nil)
//	if err != nil {
//		log.Printf("upgrade error: %v", err)
//		return
//	}
//	// 注册客户端
//	mutex.Lock()
//	clientsV2[ws] = true
//	mutex.Unlock()
//
//	defer func() {
//		mutex.Lock()
//		delete(clientsV2, ws)
//		mutex.Unlock()
//		ws.Close()
//	}()
//
//	for {
//		var message string
//		err := ws.ReadJSON(&message)
//		if err != nil {
//			log.Printf("read json error: %v", err)
//			return
//		}
//		// 广播消息（简化示例）
//		mutex.RLock()
//		for client := range clientsV2 {
//			client.WriteJSON(message)
//		}
//		mutex.RUnlock()
//	}
//}
