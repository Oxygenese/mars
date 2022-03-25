package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mars-projects/mars/app/system/internal/biz"
	"github.com/mars-projects/mars/app/system/internal/dto"
	"github.com/mars-projects/mars/app/system/internal/models"
	"github.com/mars-projects/mars/common/api"
	"github.com/mars-projects/mars/common/middleware/authentication"
)

type SysMenuHandler struct {
	api.Api
	menuBiz *biz.SysMenu
	roleBiz *biz.SysRole
}

func NewMenuHandler(menu *biz.SysMenu, role *biz.SysRole) *SysMenuHandler {
	return &SysMenuHandler{menuBiz: menu, roleBiz: role, Api: api.Api{}}
}

// GetPage Menu列表数据
// @Summary Menu列表数据
// @Description 获取JSON
// @Tags 菜单
// @Param menuName query string false "menuName"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menu [get]
// @Security Bearer
func (e SysMenuHandler) GetPage(c *gin.Context) {
	req := dto.SysMenuGetPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Query).
		Errors
	if err != nil {
		e.InternalErrorResult(err)
		return
	}
	var list = make([]models.SysMenu, 0)
	err = e.menuBiz.GetPage(&req, &list).Error
	if err != nil {
		e.ErrorResult(500, err, "查询失败")
		return
	}
	fmt.Println(list)
	e.Result(list, "查询成功")
}

// Get 获取菜单详情
// @Summary Menu详情数据
// @Description 获取JSON
// @Tags 菜单
// @Param id path string false "id"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menu/{id} [get]
// @Security Bearer
func (e SysMenuHandler) Get(c *gin.Context) {
	req := dto.SysMenuGetReq{}
	err := e.MakeContext(c).
		Bind(&req, nil).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	var object = models.SysMenu{}

	err = e.menuBiz.Get(&req, &object).Error
	if err != nil {
		e.ErrorResult(500, err, "查询失败")
		return
	}
	e.Result(object, "查询成功")
}

// Insert 创建菜单
// @Summary 创建菜单
// @Description 获取JSON
// @Tags 菜单
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysMenuInsertReq true "data"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menu [post]
// @Security Bearer
func (e SysMenuHandler) Insert(c *gin.Context) {
	req := dto.SysMenuInsertReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(c.GetInt(authentication.UserId))
	err = e.menuBiz.Insert(&req).Error
	if err != nil {
		e.ErrorResult(500, err, "创建失败")
		return
	}
	e.Result(req.GetId(), "创建成功")
}

// Update 修改菜单
// @Summary 修改菜单
// @Description 获取JSON
// @Tags 菜单
// @Accept  application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SysMenuUpdateReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menu/{id} [put]
// @Security Bearer
func (e SysMenuHandler) Update(c *gin.Context) {
	req := dto.SysMenuUpdateReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON, nil).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}

	req.SetUpdateBy(c.GetInt(authentication.UserId))
	err = e.menuBiz.Update(&req).Error
	if err != nil {
		e.ErrorResult(500, err, "更新失败")
		return
	}
	e.Result(req.GetId(), "更新成功")
}

// Delete 删除菜单
// @Summary 删除菜单
// @Description 删除数据
// @Tags 菜单
// @Param data body dto.SysMenuDeleteReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menu [delete]
// @Security Bearer
func (e SysMenuHandler) Delete(c *gin.Context) {
	control := new(dto.SysMenuDeleteReq)
	err := e.MakeContext(c).
		Bind(control, binding.JSON).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	err = e.menuBiz.Remove(control).Error
	if err != nil {
		e.ErrorResult(500, err, "删除失败")
		return
	}
	e.Result(control.GetId(), "删除成功")
}

// GetMenuRole 根据登录角色名称获取菜单列表数据（左菜单使用）
// @Summary 根据登录角色名称获取菜单列表数据（左菜单使用）
// @Description 获取JSON
// @Tags 菜单
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menurole [get]
// @Security Bearer
func (e SysMenuHandler) GetMenuRole(c *gin.Context) {
	err := e.MakeContext(c).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	result, err := e.menuBiz.SetMenuRole("admin")
	if err != nil {
		e.ErrorResult(500, err, "查询失败")
		return
	}
	e.Result(result, "")
}

//// GetMenuIDS 获取角色对应的菜单id数组
//// @Summary 获取角色对应的菜单id数组，设置角色权限使用
//// @Description 获取JSON
//// @Tags 菜单
//// @Param id path int true "id"
//// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
//// @Router /api/v1/menuids/{id} [get]
//// @Security Bearer
//func (e SysMenu) GetMenuIDS(c *gin.Context) {
//	s := new(service.SysMenu)
//	r := service.SysRole{}
//	m := dto.SysRoleByName{}
//	err := e.MakeContext(c).
//		MakeOrm().
//		Bind(&m, binding.JSON).
//		MakeService(&s.Service).
//		MakeService(&r.Service).
//		Errors
//	if err != nil {
//		e.Logger.Error(err)
//		e.Error(500, err, err.Error())
//		return
//	}
//	var data models.SysRole
//	err = r.GetWithName(&m, &data).Error
//
//	//data.RoleName = c.GetString("role")
//	//data.UpdateBy = user.GetUserId(c)
//	//result, err := data.GetIDS(s.Orm)
//
//	if err != nil {
//		e.Logger.Errorf("GetIDS error, %s", err.Error())
//		e.Error(500, err, "获取失败")
//		return
//	}
//	e.OK(result, "")
//}

// GetMenuTreeSelect 根据角色ID查询菜单下拉树结构
// @Summary 角色修改使用的菜单列表
// @Description 获取JSON
// @Tags 菜单
// @Accept  application/json
// @Product application/json
// @Param roleId path int true "roleId"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menuTreeselect/{roleId} [get]
// @Security Bearer
func (e SysMenuHandler) GetMenuTreeSelect(c *gin.Context) {
	req := dto.SelectRole{}
	err := e.MakeContext(c).
		Bind(&req, nil).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}

	result, err := e.menuBiz.SetLabel()
	if err != nil {
		e.ErrorResult(500, err, "查询失败")
		return
	}

	menuIds := make([]int, 0)
	if req.RoleId != 0 {
		menuIds, err = e.roleBiz.GetRoleMenuId(req.RoleId)
		if err != nil {
			e.ErrorResult(500, err, "")
			return
		}
	}
	e.Result(gin.H{
		"menus":       result,
		"checkedKeys": menuIds,
	}, "获取成功")
}
