package task

import (
	"github.com/google/wire"
	"github.com/mars-projects/mars/common/transaction"
)

var ProviderTasksManagerSet = wire.NewSet(NewTaskManager)

type TasksManager struct {
	*transaction.TransactionEngine
}

func NewTaskManager(engine *transaction.TransactionEngine) *TasksManager {
	var err error
	manager := &TasksManager{TransactionEngine: engine}
	err = manager.RegisterExecutor(transaction.CreateImage, false, NewCreateImageExecutor())
	err = manager.Start()
	if err != nil {
		panic(err)
	}
	return manager
}
