package task

import (
	"github.com/google/wire"
	"github.com/mars-projects/mars/app/system/internal/biz"
	"github.com/mars-projects/mars/common/transaction"
)

var ProviderTasksManagerSet = wire.NewSet(NewTaskManager)

type TasksManager struct {
	*transaction.Engine
}

func NewTaskManager(engine *transaction.Engine, option *biz.BizsOption) *TasksManager {
	var err error
	manager := &TasksManager{Engine: engine}
	err = manager.RegisterExecutor(transaction.FindSysUser, false, &FindSysUserExecutor{option.SysUser})
	err = manager.RegisterExecutor(transaction.SysUserInfo, false, &SysUserInfoExecutor{option.SysUser})
	err = manager.RegisterExecutor(transaction.CreateSysUser, false, &CreateSysUserExecutor{})

	err = manager.RegisterExecutor(transaction.QuerySysConfig, false, &QuerySysConfigExecutor{option.SysConfig})
	err = manager.RegisterExecutor(transaction.QueryAppConfig, false, &QueryAppConfigExecutor{option.SysConfig})
	err = manager.RegisterExecutor(transaction.QuerySysConfigSet, false, &QuerySysConfigSetExecutor{option.SysConfig})

	err = manager.RegisterExecutor(transaction.QuerySysMenuRole, false, &GetSysMenuRoleExecutor{option.SysMenu})

	err = manager.Start()
	if err != nil {
		panic(err)
	}
	return manager
}
