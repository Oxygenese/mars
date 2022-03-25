package task

import (
	"github.com/go-kratos/kratos/v2/log"
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
	err = manager.RegisterExecutor(transaction.CreateSysUser, false, &CreateSysUserExecutor{option.SysUser})
	err = manager.RegisterExecutor(transaction.QuerySysUserPage, false, &QuerySysUserPageExecutor{option.SysUser})
	err = manager.RegisterExecutor(transaction.ChangeSysUserStatus, false, &ChangeSysUserStatus{option.SysUser})
	err = manager.RegisterExecutor(transaction.QuerySysUserProfile, false, &QuerySysUserProfileExecutor{option.SysUser})
	err = manager.RegisterExecutor(transaction.UpdateSysUserPwd, false, &UpdateSysUserPwdExecutor{option.SysUser})
	err = manager.RegisterExecutor(transaction.ResetSysUserPwd, false, &ResetSysUserPwdExecutor{option.SysUser})
	err = manager.RegisterExecutor(transaction.UpdateSysUser, false, &UpdateSysUserExecutor{option.SysUser})
	err = manager.RegisterExecutor(transaction.DeleteSysUser, false, &DeleteSysUserExecutor{option.SysUser})
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

	// 角色模块
	err = manager.RegisterExecutor(transaction.QuerySysRolePage, false, &QuerySysRolePageExecutor{option.SysRole})
	err = manager.RegisterExecutor(transaction.QuerySysRoleById, false, &QuerySysRoleByIdExecutor{option.SysRole})
	err = manager.RegisterExecutor(transaction.UpdateSysRole, false, &UpdateSysRoleExecutor{option.SysRole})
	err = manager.RegisterExecutor(transaction.CreateSysRole, false, &CreateSysRoleExecutor{option.SysRole})
	err = manager.RegisterExecutor(transaction.ChangeSysRoleStatus, false, &ChangeSysRoleStatusExecutor{option.SysRole})
	err = manager.RegisterExecutor(transaction.DeleteSysRole, false, &DeleteSysRoleExecutor{option.SysRole})
	err = manager.RegisterExecutor(transaction.UpdateSysRoleDataScope, false, &UpdateSysRoleDataScopeExecutor{option.SysRole})

	// 菜单模块
	err = manager.RegisterExecutor(transaction.QuerySysMenuTreeSelect, false, &QuerySysMenuTreeSelectExecutor{option.SysRole, option.SysMenu})
	err = manager.RegisterExecutor(transaction.QuerySysMenuRole, false, &QuerySysMenuRoleExecutor{option.SysMenu})
	err = manager.RegisterExecutor(transaction.QuerySysMenuPage, false, &QuerySysMenuPageExecutor{option.SysMenu})
	err = manager.RegisterExecutor(transaction.QuerySysMenuById, false, &QuerySysMenuByIdExecutor{option.SysMenu})
	err = manager.RegisterExecutor(transaction.CreateSysMenu, false, &CreateSysMenuExecutor{option.SysMenu})
	err = manager.RegisterExecutor(transaction.UpdateSysMenu, false, &UpdateSysMenuExecutor{option.SysMenu})
	err = manager.RegisterExecutor(transaction.DeleteSysMenu, false, &DeleteSysMenuExecutor{option.SysMenu})

	// 数据字典模块
	err = manager.RegisterExecutor(transaction.QueryDictDataSelect, false, &QueryDictDataSelectExecutor{option.SysDictData})
	// 部门模块
	err = manager.RegisterExecutor(transaction.QuerySysDeptTree, false, &QuerySysDeptTreeExecutor{option.SysDept})
	err = manager.RegisterExecutor(transaction.QuerySysDeptTreeRoleSelect, false, &QuerySysDeptTreeRoleSelectExecutor{option.SysDept})

	// 岗位模块
	err = manager.RegisterExecutor(transaction.QuerySysPostPage, false, &QuerySysPostPageExecutor{option.SysPost})
	err = manager.Start()
	if err != nil {
		log.Errorf("[TaskManager] executor registry err :%s", err)
	}
	return manager
}
