### ChatMessage (服务端 → 客户端)

```json
{
  "type": "world | room | private | system",
  "subType": "message | user_join | user_leave | room_created",
  "fromUserId": 123,
  "fromNickname": "某某",
  "roomId": "xxx",
  "content": "xxx",
  "time": "2025-12-08T12:34:56"
}

**这里的“放在哪里”** 👉 放在你的项目文档里，用来说明“聊天协议”。

---

## 总结一句人话版

这个 JSON **不是一个要单独存起来的文件**，而是你聊天系统里「**一条消息长什么样**」的**约定**：

- 在前端：定义成一个 TypeScript interface / JSDoc 类型，用来写 WebSocket 的 `send` 和 `onmessage`。
- 在后端：定义成一个 Go struct，用来 `json.Unmarshal` / `json.Marshal`。
- 在文档：写在一个协议文档里，提醒自己和队友“所有聊天消息都按这个格式来”。

你下一步如果愿意，我可以帮你把「前端消息类型定义 + 后端 struct + 一条从前端发到后端再广播出去的完整流程图」给你画成一个“数据流思路”，方便你对着实现。
