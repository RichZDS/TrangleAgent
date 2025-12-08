package ws

import (
	"sync"

	chatmodel "leke/ws/model" // 按实际路径修改
)

// Hub 聊天中枢：管理所有连接和房间
type Hub struct {
	mu sync.RWMutex

	// 所有在线用户：userId -> Client
	clients map[uint64]*Client

	// 房间成员：roomId -> (userId -> Client)
	rooms map[string]map[uint64]*Client
}

// 全局唯一 Hub（简单起见用一个全局变量）
var ChatHub = NewHub()

// NewHub 创建一个新的 Hub 实例
func NewHub() *Hub {
	return &Hub{
		clients: make(map[uint64]*Client),
		rooms:   make(map[string]map[uint64]*Client),
	}
}

// Register 注册新连接
func (h *Hub) Register(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.clients[c.UserId] = c
	// 默认认为所有在线用户都在世界频道，这里你可以根据需要扩展
}

// Unregister 移除连接（断线/退出）
func (h *Hub) Unregister(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	delete(h.clients, c.UserId)

	// 同时把他从所有房间移除
	for roomId, members := range h.rooms {
		if _, ok := members[c.UserId]; ok {
			delete(members, c.UserId)
			if len(members) == 0 {
				delete(h.rooms, roomId)
			}
		}
	}
}

// JoinRoom 加入房间
func (h *Hub) JoinRoom(userId uint64, roomId string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	client, ok := h.clients[userId]
	if !ok {
		return
	}

	members, ok := h.rooms[roomId]
	if !ok {
		members = make(map[uint64]*Client)
		h.rooms[roomId] = members
	}
	members[userId] = client
	client.Rooms[roomId] = true
}

// LeaveRoom 离开房间
func (h *Hub) LeaveRoom(userId uint64, roomId string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	client, ok := h.clients[userId]
	if !ok {
		return
	}

	if members, ok := h.rooms[roomId]; ok {
		delete(members, userId)
		if len(members) == 0 {
			delete(h.rooms, roomId)
		}
	}
	delete(client.Rooms, roomId)
}

// BroadcastWorld 向所有在线用户发送世界频道消息
func (h *Hub) BroadcastWorld(msg *chatmodel.ChatMessage) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, client := range h.clients {
		_ = client.EnqueueMessage(msg)
	}
}

// BroadcastRoom 向房间内所有用户发送消息
func (h *Hub) BroadcastRoom(roomId string, msg *chatmodel.ChatMessage) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	members, ok := h.rooms[roomId]
	if !ok {
		return
	}
	for _, client := range members {
		_ = client.EnqueueMessage(msg)
	}
}

// SendPrivate 发送私聊消息
func (h *Hub) SendPrivate(fromUserId, toUserId uint64, msg *chatmodel.ChatMessage) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	toClient, ok := h.clients[toUserId]
	if !ok {
		return // 对方不在线，可以后面扩展：离线消息存DB
	}

	msg.FromUserId = fromUserId
	_ = toClient.EnqueueMessage(msg)
}
