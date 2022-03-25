package transaction

const (
	QueryImageList = "queryImageList"
	CreateImage    = "createImage"
)

// 用户操作
const (
	CreateSysUser       = "createSysUser"
	FindSysUser         = "findSysUser"
	QuerySysUserInfo    = "querySysUserInfo"
	QuerySysUserPage    = "querySysUserPage"
	QuerySysUserById    = "querySysUserById"
	ChangeSysUserStatus = "changeSysUserStatus"
	QuerySysUserProfile = "querySysUserProfile"
	UpdateSysUserPwd    = "updateSysUserPwd"
)

// 角色菜单操作
const (
	QuerySysMenuRole = "querySysMenuRole"
)

// 系统配置操作
const (
	QueryAppConfig      = "queryAppConfig"
	QuerySysConfigSet   = "querySysConfigSet"
	QuerySysConfigPage  = "querySysConfigPage"
	UpdateSysConfigSet  = "updateSysConfigSet"
	UpdateSysConfig     = "updateSysConfig"
	QuerySysConfigById  = "querySysConfigById"
	QuerySysConfigByKey = "querySysConfigByKey"
	CreateSysConfig     = "createSysConfig"
	DeleteSysConfig     = "deleteSysConfig"
)

// QueryDictDataSelect 数据字典操作关键字
const (
	QueryDictDataSelect = "queryDictDataSelect"
)

// 部门操作

const (
	QuerySysDeptTree = "querySysDeptTree"
)
