package task

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/mars-projects/mars/api"
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
	err = manager.RegisterExecutor(api.Operate_FindSysUser, false, &FindSysUserExecutor{option.SysUser})
	err = manager.RegisterExecutor(api.Operate_QuerySysUserInfo, false, &SysUserInfoExecutor{option.SysUser})
	err = manager.RegisterExecutor(api.Operate_QuerySysUserById, false, &QuerySysUserByIdExecutor{option.SysUser})
	err = manager.RegisterExecutor(api.Operate_CreateSysUser, false, &CreateSysUserExecutor{option.SysUser})
	err = manager.RegisterExecutor(api.Operate_QuerySysUserPage, false, &QuerySysUserPageExecutor{option.SysUser})
	err = manager.RegisterExecutor(api.Operate_ChangeSysUserStatus, false, &ChangeSysUserStatus{option.SysUser})
	err = manager.RegisterExecutor(api.Operate_QuerySysUserProfile, false, &QuerySysUserProfileExecutor{option.SysUser})
	err = manager.RegisterExecutor(api.Operate_UpdateSysUserPwd, false, &UpdateSysUserPwdExecutor{option.SysUser})
	err = manager.RegisterExecutor(api.Operate_ResetSysUserPwd, false, &ResetSysUserPwdExecutor{option.SysUser})
	err = manager.RegisterExecutor(api.Operate_UpdateSysUser, false, &UpdateSysUserExecutor{option.SysUser})
	err = manager.RegisterExecutor(api.Operate_DeleteSysUser, false, &DeleteSysUserExecutor{option.SysUser})
	// 系统设置模块
	err = manager.RegisterExecutor(api.Operate_QueryAppConfig, false, &QueryAppConfigExecutor{option.SysConfig})
	err = manager.RegisterExecutor(api.Operate_QuerySysConfigSet, false, &QuerySysConfigSetExecutor{option.SysConfig})
	err = manager.RegisterExecutor(api.Operate_QuerySysConfigPage, false, &QuerySysConfigPageExecutor{option.SysConfig})
	err = manager.RegisterExecutor(api.Operate_UpdateSysConfigSet, false, &UpdateSysConfigSetExecutor{option.SysConfig})
	err = manager.RegisterExecutor(api.Operate_QuerySysConfigById, false, &QuerySysConfigByIdExecutor{option.SysConfig})
	err = manager.RegisterExecutor(api.Operate_QuerySysConfigByKey, false, &QuerySysConfigByKeyExecutor{option.SysConfig})
	err = manager.RegisterExecutor(api.Operate_UpdateSysConfig, false, &UpdateSysConfigExecutor{option.SysConfig})
	err = manager.RegisterExecutor(api.Operate_CreateSysConfig, false, &CreateSysConfigExecutor{option.SysConfig})
	err = manager.RegisterExecutor(api.Operate_DeleteSysConfig, false, &DeleteSysConfigExecutor{option.SysConfig})
	//// 角色模块
	err = manager.RegisterExecutor(api.Operate_QuerySysRolePage, false, &QuerySysRolePageExecutor{option.SysRole})
	err = manager.RegisterExecutor(api.Operate_QuerySysRoleById, false, &QuerySysRoleByIdExecutor{option.SysRole})
	err = manager.RegisterExecutor(api.Operate_UpdateSysRole, false, &UpdateSysRoleExecutor{option.SysRole})
	err = manager.RegisterExecutor(api.Operate_CreateSysRole, false, &CreateSysRoleExecutor{option.SysRole})
	err = manager.RegisterExecutor(api.Operate_ChangeSysRoleStatus, false, &ChangeSysRoleStatusExecutor{option.SysRole})
	err = manager.RegisterExecutor(api.Operate_DeleteSysRole, false, &DeleteSysRoleExecutor{option.SysRole})
	err = manager.RegisterExecutor(api.Operate_UpdateSysRoleDataScope, false, &UpdateSysRoleDataScopeExecutor{option.SysRole})
	//// 菜单模块
	err = manager.RegisterExecutor(api.Operate_QuerySysMenuTreeSelect, false, &QuerySysMenuTreeSelectExecutor{option.SysRole, option.SysMenu})
	err = manager.RegisterExecutor(api.Operate_QuerySysMenuRole, false, &QuerySysMenuRoleExecutor{option.SysMenu})
	err = manager.RegisterExecutor(api.Operate_QuerySysMenuPage, false, &QuerySysMenuPageExecutor{option.SysMenu})
	err = manager.RegisterExecutor(api.Operate_QuerySysMenuById, false, &QuerySysMenuByIdExecutor{option.SysMenu})
	err = manager.RegisterExecutor(api.Operate_CreateSysMenu, false, &CreateSysMenuExecutor{option.SysMenu})
	err = manager.RegisterExecutor(api.Operate_UpdateSysMenu, false, &UpdateSysMenuExecutor{option.SysMenu})
	err = manager.RegisterExecutor(api.Operate_DeleteSysMenu, false, &DeleteSysMenuExecutor{option.SysMenu})

	// 数据字典模块
	err = manager.RegisterExecutor(api.Operate_QueryDictDataSelect, false, &QueryDictDataSelectExecutor{option.SysDictData})
	// 部门模块
	err = manager.RegisterExecutor(api.Operate_QuerySysDeptTree, false, &QuerySysDeptTreeExecutor{option.SysDept})
	err = manager.RegisterExecutor(api.Operate_QuerySysDeptTreeRoleSelect, false, &QuerySysDeptTreeRoleSelectExecutor{option.SysDept})

	// 岗位模块
	err = manager.RegisterExecutor(api.Operate_QuerySysPostPage, false, &QuerySysPostPageExecutor{option.SysPost})
	err = manager.Start()
	if err != nil {
		log.Errorf("[TaskManager] executor registry err :%s", err)
	}
	return manager
}
