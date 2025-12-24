package ws

import (
	"encoding/json"

	chatmodel "leke/ws/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

// Client 表示一个在线用户的 WebSocket 连接
// 这里只是框架：先把字段和发送逻辑搭好，读写循环可以后面慢慢加。
type Client struct {
	UserId   uint64         // 当前用户ID
	Nickname string         // 昵称（可选）
	Conn     *ghttp.Request // 底层 WebSocket 连接

	// Send 是一个发送队列：其他地方往这里丢 []byte，这个 client 的写协程负责发出去
	Send chan []byte

	// Rooms 记录当前用户加入的房间（简单用 map 做集合）
	Rooms map[string]bool
}

// NewClient 创建一个新的客户端连接对象
func NewClient(userId uint64, nickname string, conn *ghttp.Request) *Client {
	return &Client{
		UserId:   userId,
		Nickname: nickname,
		Conn:     conn,
		Send:     make(chan []byte, 256), // 简单先给一个缓冲，防止轻微阻塞
		Rooms:    make(map[string]bool),
	}
}

// EnqueueMessage 将 ChatMessage 编码为 JSON 丢到 Send 队列
// 真正 WriteMessage 的动作建议在一个单独的 writeLoop 里做。
func (c *Client) EnqueueMessage(msg *chatmodel.ChatMessage) error {
	// 如果需要，补一些默认字段
	if msg.FromUserId == 0 {
		msg.FromUserId = c.UserId
	}
	if msg.FromNickname == "" {
		msg.FromNickname = c.Nickname
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	select {
	case c.Send <- data:
		// 正常入队
	default:
		// 队列满了，可以考虑：丢弃 / 断开连接 / 打日志
		// 这里简单选择丢弃，并返回错误
		return ErrSendQueueFull
	}
	return nil
}
