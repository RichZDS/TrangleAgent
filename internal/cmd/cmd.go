package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"leke/internal/controller"
	"leke/internal/controller/websocket"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// ... existing code ...
			// CORS 中间件
			s.Use(func(r *ghttp.Request) {
				r.Response.CORS(ghttp.CORSOptions{
					AllowOrigin:      "http://localhost:3000,http://localhost:8080",
					AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
					AllowHeaders:     "Content-Type,Authorization,X-Requested-With",
					AllowCredentials: "true",
					MaxAge:           3600,
				})
				if r.Method == "OPTIONS" {
					r.Response.WriteStatusExit(200)
					return
				}
				r.Middleware.Next()
			})
			// ... existing code ...

			// 注册 API 路由组
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				controller.RegisterControllers(group)
			})

			// 注册 WebSocket 聊天室路由（不在 /api 组下，因为 WebSocket 不需要中间件）
			s.BindHandler("/ws/chat", websocket.HandleChatConnections)

			// 配置静态文件服务（用于提供 HTML 客户端页面）
			s.SetServerRoot("resource/public")

			s.Run()
			return nil
		},
	}
)
