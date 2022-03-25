package transaction

import (
	"github.com/google/wire"
	"github.com/mars-projects/mars/common/transaction"
)

var ProviderTransactionEngineSet = wire.NewSet(NewTransactionEngine)

func NewTransactionEngine(sender transaction.Sender) *transaction.Engine {
	engine, err := transaction.NewTransactionEngine(1<<20, sender)
	if err != nil {
		panic(err)
	}
	return engine
}
