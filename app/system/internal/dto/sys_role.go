package dto

import (
	"github.com/mars-projects/mars/app/system/internal/models"
	"github.com/mars-projects/mars/common/dto"
	common "github.com/mars-projects/mars/common/models"
)

type SysRoleGetPageReq struct {
	dto.Pagination `search:"-"`
	SysRoleOrder
	RoleId    int    `json:"roleId" search:"type:exact;column:role_id;table:sys_role" comment:"角色编码"`        // 角色编码
	RoleName  string `json:"roleName" search:"type:contains;column:role_name;table:sys_role" comment:"角色名称"` // 角色名称
	Status    string `json:"status" search:"type:exact;column:status;table:sys_role" comment:"状态"`           // 状态
	RoleKey   string `json:"roleKey" search:"type:contains;column:role_key;table:sys_role" comment:"角色代码"`   // 角色代码
	RoleSort  int    `json:"roleSort" search:"type:exact;column:role_sort;table:sys_role" comment:"角色排序"`    // 角色排序
	Flag      string `json:"flag" search:"type:exact;column:flag;table:sys_role" comment:"标记"`               // 标记
	Remark    string `json:"remark" search:"type:exact;column:remark;table:sys_role" comment:"备注"`           // 备注
	Admin     bool   `json:"admin" search:"type:exact;column:admin;table:sys_role" comment:"是否管理员"`
	DataScope string `json:"dataScope" search:"type:exact;column:data_scope;table:sys_role" comment:"是否管理员"`
}

type SysRoleOrder struct {
	RoleIdOrder    string `search:"type:order;column:role_id;table:sys_role" json:"roleIdOrder"`
	RoleNameOrder  string `search:"type:order;column:role_name;table:sys_role" json:"roleNameOrder"`
	RoleSortOrder  string `search:"type:order;column:role_sort;table:sys_role" json:"usernameOrder"`
	StatusOrder    string `search:"type:order;column:status;table:sys_role" json:"statusOrder"`
	CreatedAtOrder string `search:"type:order;column:created_at;table:sys_role" json:"createdAtOrder"`
}

func (m *SysRoleGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysRoleInsertReq struct {
	RoleId    int              `json:"id" comment:"角色编码"`       // 角色编码
	RoleName  string           `json:"roleName" comment:"角色名称"` // 角色名称
	Status    string           `json:"status" comment:"状态"`     // 状态
	RoleKey   string           `json:"roleKey" comment:"角色代码"`  // 角色代码
	RoleSort  int              `json:"roleSort" comment:"角色排序"` // 角色排序
	Flag      string           `json:"flag" comment:"标记"`       // 标记
	Remark    string           `json:"remark" comment:"备注"`     // 备注
	Admin     bool             `json:"admin" comment:"是否管理员"`
	DataScope string           `json:"dataScope"`
	SysMenu   []models.SysMenu `json:"sysMenu"`
	MenuIds   []int            `json:"menuIds"`
	SysDept   []models.SysDept `json:"sysDept"`
	DeptIds   []int            `json:"deptIds"`
	common.ControlBy
}

func (s *SysRoleInsertReq) Generate(model *models.SysRole) {
	if s.RoleId != 0 {
		model.RoleId = s.RoleId
	}
	model.RoleName = s.RoleName
	model.Status = s.Status
	model.RoleKey = s.RoleKey
	model.RoleSort = s.RoleSort
	model.Flag = s.Flag
	model.Remark = s.Remark
	model.Admin = s.Admin
	model.DataScope = s.DataScope
	model.SysMenu = &s.SysMenu
	model.SysDept = s.SysDept
}

func (s *SysRoleInsertReq) GetId() interface{} {
	return s.RoleId
}

type SysRoleUpdateReq struct {
	RoleId    int              `uri:"id" comment:"角色编码"`        // 角色编码
	RoleName  string           `json:"roleName" comment:"角色名称"` // 角色名称
	Status    string           `json:"status" comment:"状态"`     // 状态
	RoleKey   string           `json:"roleKey" comment:"角色代码"`  // 角色代码
	RoleSort  int              `json:"roleSort" comment:"角色排序"` // 角色排序
	Flag      string           `json:"flag" comment:"标记"`       // 标记
	Remark    string           `json:"remark" comment:"备注"`     // 备注
	Admin     bool             `json:"admin" comment:"是否管理员"`
	DataScope string           `json:"dataScope"`
	SysMenu   []models.SysMenu `json:"sysMenu"`
	MenuIds   []int            `json:"menuIds"`
	SysDept   []models.SysDept `json:"sysDept"`
	DeptIds   []int            `json:"deptIds"`
	common.ControlBy
}

func (s *SysRoleUpdateReq) Generate(model *models.SysRole) {
	if s.RoleId != 0 {
		model.RoleId = s.RoleId
	}
	model.RoleName = s.RoleName
	model.Status = s.Status
	model.RoleKey = s.RoleKey
	model.RoleSort = s.RoleSort
	model.Flag = s.Flag
	model.Remark = s.Remark
	model.Admin = s.Admin
	model.DataScope = s.DataScope
	model.SysMenu = &s.SysMenu
	model.SysDept = s.SysDept
}

func (s *SysRoleUpdateReq) GetId() interface{} {
	return s.RoleId
}

type UpdateStatusReq struct {
	RoleId int    `json:"roleId" comment:"角色编码"` // 角色编码
	Status string `json:"status" comment:"状态"`   // 状态
	common.ControlBy
}

func (s *UpdateStatusReq) Generate(model *models.SysRole) {
	if s.RoleId != 0 {
		model.RoleId = s.RoleId
	}
	model.Status = s.Status
}

func (s *UpdateStatusReq) GetId() interface{} {
	return s.RoleId
}

type SysRoleByName struct {
	RoleName string `json:"role"` // 角色编码
}

type SysRoleGetReq struct {
	Id int `json:"id"`
}

func (s *SysRoleGetReq) GetId() interface{} {
	return s.Id
}

type SysRoleDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysRoleDeleteReq) GetId() interface{} {
	return s.Ids
}

// RoleDataScopeReq 角色数据权限修改
type RoleDataScopeReq struct {
	RoleId    int    `json:"roleId" binding:"required"`
	DataScope string `json:"dataScope" binding:"required"`
	DeptIds   []int  `json:"deptIds"`
}

func (s *RoleDataScopeReq) Generate(model *models.SysRole) {
	if s.RoleId != 0 {
		model.RoleId = s.RoleId
	}
	model.DataScope = s.DataScope
	model.DeptIds = s.DeptIds
}

type DeptIdList struct {
	DeptId int `json:"DeptId"`
}
