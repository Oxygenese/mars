package sender

import (
	"fmt"
	"github.com/mars-projects/mars/common/framework"
)

func NewMessageSender() framework.MessageSender {
	return &MessageSender{}
}

type MessageSender struct {
}

func (m MessageSender) SendMessage(msg framework.Message, target string) error {
	fmt.Println("执行成功:", msg)
	return nil
}

func (m MessageSender) SendToSelf(msg framework.Message) error {
	fmt.Println("消息回路：", msg)
	return nil
}
