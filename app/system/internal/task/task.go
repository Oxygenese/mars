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

//NewTaskManager 创建任务管理器 并向内注册任务执行器
func NewTaskManager(engine *transaction.Engine, option *biz.BizsOption) *TasksManager {
	var err error
	manager := &TasksManager{Engine: engine}
	// 用户模块
	err = manager.RegisterExecutor(transaction.FindSysUser, false, &FindSysUserExecutor{option.SysUser})
	err = manager.RegisterExecutor(transaction.QuerySysUserInfo, false, &SysUserInfoExecutor{option.SysUser})
	err = manager.RegisterExecutor(transaction.QuerySysUserById, false, &QuerySysUserByIdExecutor{option.SysUser})
	err = manager.RegisterExecutor(transaction.CreateSysUser, false, &CreateSysUserExecutor{})
	err = manager.RegisterExecutor(transaction.QuerySysUserPage, false, &QuerySysUserPageExecutor{option.SysUser})
	err = manager.RegisterExecutor(transaction.ChangeSysUserStatus, false, &ChangeSysUserStatus{option.SysUser})
	err = manager.RegisterExecutor(transaction.QuerySysUserProfile, false, &QuerySysUserProfileExecutor{option.SysUser})
	err = manager.RegisterExecutor(transaction.UpdateSysUserPwd, false, &UpdateSysUserPwdExecutor{option.SysUser})

	// 系统设置模块
	err = manager.RegisterExecutor(transaction.QueryAppConfig, false, &QueryAppConfigExecutor{option.SysConfig})
	err = manager.RegisterExecutor(transaction.QuerySysConfigSet, false, &QuerySysConfigSetExecutor{option.SysConfig})
	err = manager.RegisterExecutor(transaction.QuerySysConfigPage, false, &QuerySysConfigPageExecutor{option.SysConfig})
	err = manager.RegisterExecutor(transaction.UpdateSysConfigSet, false, &UpdateSysConfigSetExecutor{option.SysConfig})
	err = manager.RegisterExecutor(transaction.QuerySysConfigById, false, &QuerySysConfigByIdExecutor{option.SysConfig})
	err = manager.RegisterExecutor(transaction.QuerySysConfigByKey, false, &QuerySysConfigByKeyExecutor{option.SysConfig})
	err = manager.RegisterExecutor(transaction.UpdateSysConfig, false, &UpdateSysConfigExecutor{option.SysConfig})
	err = manager.RegisterExecutor(transaction.CreateSysConfig, false, &CreateSysConfigExecutor{option.SysConfig})
	err = manager.RegisterExecutor(transaction.DeleteSysConfig, false, &DeleteSysConfigExecutor{option.SysConfig})

	// 角色权限菜单模块
	err = manager.RegisterExecutor(transaction.QuerySysMenuRole, false, &GetSysMenuRoleExecutor{option.SysMenu})

	// 数据字典模块
	err = manager.RegisterExecutor(transaction.QueryDictDataSelect, false, &QueryDictDataSelectExecutor{option.SysDictData})
	// 部门模块
	err = manager.RegisterExecutor(transaction.QuerySysDeptTree, false, &QuerySysDeptTreeExecutor{option.SysDept})
	err = manager.Start()
	if err != nil {
		panic(err)
	}
	return manager
}
