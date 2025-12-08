package model

// MessageType 消息的大类：发到哪里
type MessageType string

const (
	TypeWorld   MessageType = "world"   // 世界频道
	TypeRoom    MessageType = "room"    // 房间
	TypePrivate MessageType = "private" // 私聊
	TypeSystem  MessageType = "system"  // 系统消息（如通知、提示等）
)

// MessageAction 消息的动作：干什么事
type MessageAction string

const (
	ActionSend       MessageAction = "send"        // 发送聊天消息
	ActionJoinRoom   MessageAction = "join_room"   // 加入房间
	ActionLeaveRoom  MessageAction = "leave_room"  // 离开房间
	ActionCreateRoom MessageAction = "create_room" // 创建房间（也可以走 HTTP）
)

// ChatMessage WebSocket 收/发的统一结构
// 前端和后端都按这个结构来编码/解码 JSON。
type ChatMessage struct {
	Type         MessageType   `json:"type"`                   // 消息类型：world/room/private/system
	Action       MessageAction `json:"action"`                 // 动作：send/join_room/leave_room...
	FromUserId   uint64        `json:"fromUserId,omitempty"`   // 发送者用户ID
	FromNickname string        `json:"fromNickname,omitempty"` // 发送者昵称
	RoomId       string        `json:"roomId,omitempty"`       // 房间ID（房间消息时使用）
	ToUserId     uint64        `json:"toUserId,omitempty"`     // 目标用户ID（私聊时使用）
	Content      string        `json:"content,omitempty"`      // 文本内容
	Time         string        `json:"time,omitempty"`         // 时间（ISO字符串，前期用 string 即可）
}
