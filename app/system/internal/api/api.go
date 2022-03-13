package api

import (
	"github.com/google/wire"
	"github.com/mars-projects/mars/app/system/internal/biz"
)

var ProviderApiOptionSet = wire.NewSet(NewApiOptions)

type ApiOption struct {
	SysUserHandler     *SysUserHandler
	SysConfigHandler   *SysConfigHandler
	SysDictTypeHandler *SysDictTypeHandler
	SysMenuHandler     *SysMenuHandler
	SysDictDataHandler *SysDictDataHandler
	SysDeptHandler     *SysDeptHandler
	SysPostHandler     *SysPostHandler
	SysRoleHandler     *SysRoleHandler
}

func NewApiOptions(option *biz.BizsOption) *ApiOption {
	return &ApiOption{
		SysUserHandler:     NewSysUserHandler(option.SysUser),
		SysConfigHandler:   NewSysConfigHandler(option.SysConfig),
		SysDictTypeHandler: NewSysDictTypHandler(option.SysDictType),
		SysMenuHandler:     NewMenuHandler(option.SysMenu, option.SysRole),
		SysDictDataHandler: NewSysDictDataHandler(option.SysDictData),
		SysDeptHandler:     NewSysDeptHandler(option.SysDept),
		SysPostHandler:     NewSysPostHandler(option.SysPost),
		SysRoleHandler:     NewSysRoleHandler(option.SysRole),
	}
}
