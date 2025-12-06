package websocket

import (
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gorilla/websocket"
)

// WebSocketController WebSocket 控制器
type WebSocketController struct{}

// New 创建 WebSocket 控制器实例
func New() *WebSocketController {
	return &WebSocketController{}
}

// wsUpgrader WebSocket 升级器配置
var wsUpgrader = websocket.Upgrader{
	// CheckOrigin 允许任何来源（开发环境）
	// 生产环境中应该实现适当的来源检查以确保安全
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	// Error 处理升级失败的错误
	Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
		// 在这里实现错误处理逻辑
		g.Log().Errorf(r.Context(), "WebSocket upgrade error: %v", reason)
	},
}

// Echo WebSocket Echo 服务器处理器
// 路径: /ws
func (c *WebSocketController) Echo(r *ghttp.Request) {
	// 将 HTTP 连接升级为 WebSocket
	ws, err := wsUpgrader.Upgrade(r.Response.Writer, r.Request, nil)
	if err != nil {
		r.Response.Write(err.Error())
		return
	}
	defer ws.Close()

	// 获取请求上下文用于日志记录
	var ctx = r.Context()
	logger := g.Log()

	// 消息处理循环
	for {
		// 读取传入的 WebSocket 消息
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			break // 连接关闭或发生错误
		}
		// 记录接收到的消息
		logger.Infof(ctx, "received message: %s", msg)
		// 将消息回显给客户端
		if err = ws.WriteMessage(msgType, msg); err != nil {
			break // 写入消息时出错
		}
	}
	// 记录连接关闭
	logger.Info(ctx, "websocket connection closed")
}
