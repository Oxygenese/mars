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
	user := &SysUserExecutor{option.SysUser}
	err = manager.RegisterExecutor(api.Operate_FindSysUser, false, user.FindSysUser)
	err = manager.RegisterExecutor(api.Operate_QuerySysUserInfo, false, user.SysUserInfo)
	err = manager.RegisterExecutor(api.Operate_QuerySysUserById, false, user.QuerySysUserById)
	err = manager.RegisterExecutor(api.Operate_CreateSysUser, false, user.CreateSysUser)
	err = manager.RegisterExecutor(api.Operate_QuerySysUserPage, false, user.QuerySysUserPage)
	err = manager.RegisterExecutor(api.Operate_ChangeSysUserStatus, false, user.ChangeSysUserStatus)
	err = manager.RegisterExecutor(api.Operate_QuerySysUserProfile, false, user.QuerySysUserProfile)
	err = manager.RegisterExecutor(api.Operate_UpdateSysUserPwd, false, user.UpdateSysUserPwd)
	err = manager.RegisterExecutor(api.Operate_ResetSysUserPwd, false, user.ResetSysUserPwd)
	err = manager.RegisterExecutor(api.Operate_UpdateSysUser, false, user.UpdateSysUser)
	err = manager.RegisterExecutor(api.Operate_DeleteSysUser, false, user.DeleteSysUser)
	// 系统设置模块
	config := &SysConfigExecutor{option.SysConfig}
	err = manager.RegisterExecutor(api.Operate_QueryAppConfig, false, config.QueryAppConfig)
	err = manager.RegisterExecutor(api.Operate_QuerySysConfigSet, false, config.QuerySysConfigSet)
	err = manager.RegisterExecutor(api.Operate_QuerySysConfigPage, false, config.QuerySysConfigPage)
	err = manager.RegisterExecutor(api.Operate_UpdateSysConfigSet, false, config.UpdateSysConfigSet)
	err = manager.RegisterExecutor(api.Operate_QuerySysConfigById, false, config.QuerySysConfigById)
	err = manager.RegisterExecutor(api.Operate_QuerySysConfigByKey, false, config.QuerySysConfigByKey)
	err = manager.RegisterExecutor(api.Operate_UpdateSysConfig, false, config.UpdateSysConfig)
	err = manager.RegisterExecutor(api.Operate_CreateSysConfig, false, config.CreateSysConfig)
	err = manager.RegisterExecutor(api.Operate_DeleteSysConfig, false, config.DeleteSysConfig)
	//// 角色模块
	role := &SysRoleExecutor{option.SysRole}
	err = manager.RegisterExecutor(api.Operate_QuerySysRolePage, false, role.QuerySysRolePage)
	err = manager.RegisterExecutor(api.Operate_QuerySysRoleById, false, role.QuerySysRoleById)
	err = manager.RegisterExecutor(api.Operate_UpdateSysRole, false, role.UpdateSysRole)
	err = manager.RegisterExecutor(api.Operate_CreateSysRole, false, role.CreateSysRole)
	err = manager.RegisterExecutor(api.Operate_ChangeSysRoleStatus, false, role.ChangeSysRoleStatus)
	err = manager.RegisterExecutor(api.Operate_DeleteSysRole, false, role.DeleteSysRole)
	err = manager.RegisterExecutor(api.Operate_UpdateSysRoleDataScope, false, role.UpdateSysRoleDataScope)
	//// 菜单模块
	menu := &SysMenuExecutor{option.SysRole, option.SysMenu}
	err = manager.RegisterExecutor(api.Operate_QuerySysMenuTreeSelect, false, menu.QuerySysMenuTreeSelect)
	err = manager.RegisterExecutor(api.Operate_QuerySysMenuRole, false, menu.QuerySysMenuRole)
	err = manager.RegisterExecutor(api.Operate_QuerySysMenuPage, false, menu.QuerySysMenuPage)
	err = manager.RegisterExecutor(api.Operate_QuerySysMenuById, false, menu.QuerySysMenuById)
	err = manager.RegisterExecutor(api.Operate_CreateSysMenu, false, menu.CreateSysMenu)
	err = manager.RegisterExecutor(api.Operate_UpdateSysMenu, false, menu.UpdateSysMenu)
	err = manager.RegisterExecutor(api.Operate_DeleteSysMenu, false, menu.DeleteSysMenu)
	// 字典数据模块
	dictData := &SysDictDataExecutor{option.SysDictData}
	err = manager.RegisterExecutor(api.Operate_QueryDictDataSelect, false, dictData.QueryDictDataSelect)
	err = manager.RegisterExecutor(api.Operate_QueryDictDataByCode, false, dictData.QueryDictDataByCode)
	err = manager.RegisterExecutor(api.Operate_QueryDictDataPage, false, dictData.QueryDictDataPage)
	err = manager.RegisterExecutor(api.Operate_CreateDictData, false, dictData.CreateDictData)
	err = manager.RegisterExecutor(api.Operate_UpdateDictData, false, dictData.UpdateDictData)
	err = manager.RegisterExecutor(api.Operate_DeleteDictData, false, dictData.DeleteDictData)
	// 字典类型模块
	dictType := &SysDictTypeExecutor{option.SysDictType}
	err = manager.RegisterExecutor(api.Operate_QueryDictTypePage, false, dictType.QueryDictTypePage)
	err = manager.RegisterExecutor(api.Operate_QueryDictTypeById, false, dictType.QueryDictTypeById)
	err = manager.RegisterExecutor(api.Operate_CreateDictType, false, dictType.CreateDictType)
	err = manager.RegisterExecutor(api.Operate_UpdateDictType, false, dictType.UpdateDictType)
	err = manager.RegisterExecutor(api.Operate_DeleteDictType, false, dictType.DeleteDictType)
	err = manager.RegisterExecutor(api.Operate_ExportDictType, false, dictType.ExportDictType)
	err = manager.RegisterExecutor(api.Operate_QueryDictTypeOptionSelect, false, dictData.QueryDictDataSelect)
	// 部门模块
	dept := &SysDeptExecutor{option.SysDept}
	err = manager.RegisterExecutor(api.Operate_QuerySysDeptTree, false, dept.QuerySysDeptTree)
	err = manager.RegisterExecutor(api.Operate_QuerySysDeptTreeRoleSelect, false, dept.QuerySysDeptTreeRoleSelect)
	err = manager.RegisterExecutor(api.Operate_QuerySysDeptById, false, dept.QuerySysDeptById)
	err = manager.RegisterExecutor(api.Operate_QuerySysDeptPage, false, dept.QuerySysDeptPage)
	err = manager.RegisterExecutor(api.Operate_DeleteSysDept, false, dept.DeleteSysDept)
	err = manager.RegisterExecutor(api.Operate_UpdateSysDept, false, dept.UpdateSysDept)

	// 岗位模块
	post := &SysPostExecutor{option.SysPost}
	err = manager.RegisterExecutor(api.Operate_QuerySysPostPage, false, post.QuerySysPostPage)
	err = manager.RegisterExecutor(api.Operate_QuerySysPostById, false, post.QuerySysPostById)
	err = manager.RegisterExecutor(api.Operate_CreateSysPost, false, post.CreateSysPost)
	err = manager.RegisterExecutor(api.Operate_UpdateSysPost, false, post.UpdateSysPost)
	err = manager.RegisterExecutor(api.Operate_DeleteSysPost, false, post.DeleteSysPost)
	err = manager.Start()
	if err != nil {
		log.Errorf("[TaskManager] executor registry err :%s", err)
	}
	return manager
}
