package transaction

import (
	"github.com/google/wire"
	"github.com/mars-projects/mars/common/framework"
)

var ProviderTransactionEngineSet = wire.NewSet(NewTransactionEngine)

func NewTransactionEngine() *framework.TransactionEngine {
	engine, err := framework.NewTransactionEngine(1 << 20)
	if err != nil {
		panic(err)
	}
	return engine
}
