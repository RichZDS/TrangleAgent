# WebSocket Echo Server 使用说明

## 功能说明

这是一个基于 GoFrame 框架实现的 WebSocket Echo 服务器示例。

## 使用方法

1. 启动服务器：
   ```bash
   go run main.go
   ```

2. 访问测试页面：
   打开浏览器访问：`http://localhost:8000/`

3. WebSocket 连接地址：
   - WebSocket: `ws://localhost:8000/ws`
   - 如果使用 HTTPS: `wss://localhost:8000/ws`

## 功能特性

- ✅ WebSocket 连接管理
- ✅ 消息回显（Echo）
- ✅ 连接状态监控
- ✅ 自动重连提示
- ✅ 支持文本消息

## 测试步骤

1. 打开浏览器访问 `http://localhost:8000/`
2. 页面会自动连接到 WebSocket 服务器
3. 在输入框中输入消息并点击"发送"按钮
4. 服务器会将消息回显到页面上

