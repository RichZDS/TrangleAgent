package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gorilla/websocket"
)

type ChatMessage struct {
	RoomId  string `json:"roomId"`
	UserId  string `json:"userId"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
var clients = make(map[*websocket.Conn]bool)
var roomClients = make(map[string]map[*websocket.Conn]bool)
var broadcast = make(chan ChatMessage, 256)
var mutex = sync.RWMutex{}
var once sync.Once

func HandleChatConnections(r *ghttp.Request) {
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
		for roomId, members := range roomClients {
			if _, ok := members[ws]; ok {
				delete(members, ws)
				if len(members) == 0 {
					delete(roomClients, roomId)
				}
			}
		}
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

		switch msg.Type {
		case "join":
			mutex.Lock()
			members, ok := roomClients[msg.RoomId]
			if !ok {
				members = make(map[*websocket.Conn]bool)
				roomClients[msg.RoomId] = members
			}
			members[ws] = true
			mutex.Unlock()
		case "leave":
			mutex.Lock()
			if members, ok := roomClients[msg.RoomId]; ok {
				delete(members, ws)
				if len(members) == 0 {
					delete(roomClients, msg.RoomId)
				}
			}
			mutex.Unlock()
		}

		select {
		case broadcast <- msg:
		default:
			log.Printf("broadcast channel full, dropping message")
		}
	}
}

func handleBroadcast() {
	for {
		msg := <-broadcast
		mutex.RLock()
		members, ok := roomClients[msg.RoomId]
		if !ok {
			mutex.RUnlock()
			continue
		}
		for client := range members {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("write error: %v", err)
				client.Close()
				mutex.RUnlock()
				mutex.Lock()
				delete(members, client)
				if len(members) == 0 {
					delete(roomClients, msg.RoomId)
				}
				mutex.Unlock()
				mutex.RLock()
			}
		}
		mutex.RUnlock()
	}
}
