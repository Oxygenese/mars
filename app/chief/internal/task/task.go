package task

import (
	"fmt"
	"github.com/ceph/go-ceph/rados"
	"github.com/google/wire"
	"github.com/mars-projects/mars/common/transaction"
)

var ProviderTasksManagerSet = wire.NewSet(NewTaskManager)

type TasksManager struct {
	*transaction.Engine
}

func NewTaskManager(engine *transaction.Engine, conn *rados.Conn) *TasksManager {
	var err error
	manager := &TasksManager{Engine: engine}
	err = manager.Start()
	pools, err := conn.ListPools()
	if err != nil {
		return nil
	}
	fmt.Println("存储池：", pools)
	if err != nil {
		panic(err)
	}
	return manager
}
