package ws

import "errors"

var (
	ErrSendQueueFull = errors.New("客户端队列堵塞")
)
