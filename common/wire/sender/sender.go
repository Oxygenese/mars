package sender

import (
	"github.com/google/wire"
	"github.com/mars-projects/mars/common/framework"
	"github.com/mars-projects/mars/common/framework/sender"
)

var ProviderMessageSenderSet = wire.NewSet(NewMessageSender)

func NewMessageSender() framework.MessageSender {
	return sender.NewMessageSender()
}
