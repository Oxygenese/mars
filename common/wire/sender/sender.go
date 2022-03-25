package sender

import (
	"github.com/google/wire"
	"github.com/mars-projects/mars/common/transaction"
	"github.com/mars-projects/mars/common/ws"
)

var ProviderMessageSenderSet = wire.NewSet(NewMessageSender)

func NewMessageSender() transaction.Sender {
	return ws.NewWs()
}
