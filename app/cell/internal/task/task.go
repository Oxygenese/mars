package task

import (
	"github.com/google/wire"
	"github.com/mars-projects/mars/common/transaction"
)

var ProviderTasksManagerSet = wire.NewSet(NewTaskManager)

type TasksManager struct {
	*transaction.Engine
}

func NewTaskManager(engine *transaction.Engine) *TasksManager {
	var err error
	manager := &TasksManager{Engine: engine}
	err = manager.Start()
	if err != nil {
		panic(err)
	}
	return manager
}
