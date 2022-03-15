package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mars-projects/mars/app/system/internal/biz"
	"github.com/mars-projects/mars/app/system/internal/dto"
	"github.com/mars-projects/mars/app/system/internal/models"
	"github.com/mars-projects/mars/lib/api"
	"github.com/mars-projects/mars/lib/wire/middleware/oauth"
	"net/http"
)

type SysRoleHandler struct {
	api.Api
	biz *biz.SysRole
}

func NewSysRoleHandler(role *biz.SysRole) *SysRoleHandler {
	return &SysRoleHandler{
		Api: api.Api{},
		biz: role,
	}
}

// GetPage
// @Summary 角色列表数据
// @Description Get JSON
// @Tags 角色/Role
// @Param roleName query string false "roleName"
// @Param status query string false "status"
// @Param roleKey query string false "roleKey"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/role [get]
// @Security Bearer
func (e SysRoleHandler) GetPage(c *gin.Context) {
	req := dto.SysRoleGetPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}

	list := make([]models.SysRole, 0)
	var count int64

	err = e.biz.GetPage(&req, &list, &count)
	if err != nil {
		e.ErrorResult(500, err, "查询失败")
		return
	}

	e.PageResult(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get
// @Summary 获取Role数据
// @Description 获取JSON
// @Tags 角色/Role
// @Param roleId path string false "roleId"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/role/{id} [get]
// @Security Bearer
func (e SysRoleHandler) Get(c *gin.Context) {
	req := dto.SysRoleGetReq{}
	err := e.MakeContext(c).
		Bind(&req, nil).
		Errors
	if err != nil {
		e.ErrorResult(500, err, fmt.Sprintf(" %s ", err.Error()))
		return
	}

	var object models.SysRole

	err = e.biz.Get(&req, &object)
	if err != nil {
		e.ErrorResult(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}

	e.Result(object, "查询成功")
}

// Insert
// @Summary 创建角色
// @Description 获取JSON
// @Tags 角色/Role
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysRoleInsertReq true "data"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/role [post]
// @Security Bearer
func (e SysRoleHandler) Insert(c *gin.Context) {
	req := dto.SysRoleInsertReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}

	// 设置创建人
	req.CreateBy = oauth.GetUserId(c)
	if req.Status == "" {
		req.Status = "2"
	}
	err = e.biz.Insert(&req)
	if err != nil {
		e.ErrorResult(500, err, "创建失败")
		return
	}
	e.Result(req.GetId(), "创建成功")
}

// Update 修改用户角色
// @Summary 修改用户角色
// @Description 获取JSON
// @Tags 角色/Role
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysRoleUpdateReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/role/{id} [put]
// @Security Bearer
func (e SysRoleHandler) Update(c *gin.Context) {
	req := dto.SysRoleUpdateReq{}
	err := e.MakeContext(c).
		Bind(&req, nil, binding.JSON).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	//cb := sdk.Runtime.GetCasbinKey(c.Request.Host)

	req.SetUpdateBy(oauth.GetUserId(c))

	err = e.biz.Update(&req)
	if err != nil {
		return
	}
	//_, err = global.LoadPolicy(c)
	//if err != nil {
	//	e.Error(500, err, "")
	//	return
	//}
	e.Result(req.GetId(), "更新成功")
}

// Delete
// @Summary 删除用户角色
// @Description 删除数据
// @Tags 角色/Role
// @Param data body dto.SysRoleDeleteReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/role [delete]
// @Security Bearer
func (e SysRoleHandler) Delete(c *gin.Context) {
	req := dto.SysRoleDeleteReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON).
		Errors
	if err != nil {
		e.ErrorResult(500, err, fmt.Sprintf("删除角色 %v 失败，\r\n失败信息 %s", req.Ids, err.Error()))
		return
	}

	err = e.biz.Remove(&req)
	if err != nil {
		e.ErrorResult(500, err, "")
		return
	}
	//_, err = global.LoadPolicy(c)
	//if err != nil {
	//	e.Error(500, err, fmt.Sprintf("删除角色 %v 失败，失败信息 %s", req.GetId(), err.Error()))
	//	return
	//}
	e.Result(req.GetId(), fmt.Sprintf("删除角色角色 %v 状态成功！", req.GetId()))
}

// Update2Status 修改用户角色状态
// @Summary 修改用户角色
// @Description 获取JSON
// @Tags 角色/Role
// @Accept  application/json
// @Product application/json
// @Param data body dto.UpdateStatusReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/role-status/{id} [put]
// @Security Bearer
func (e SysRoleHandler) Update2Status(c *gin.Context) {
	req := dto.UpdateStatusReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON, nil).
		Errors
	if err != nil {
		e.ErrorResult(500, err, fmt.Sprintf("更新角色状态失败，失败原因：%s ", err.Error()))
		return
	}
	req.SetUpdateBy(oauth.GetUserId(c))
	err = e.biz.UpdateStatus(&req)
	if err != nil {
		e.ErrorResult(500, err, fmt.Sprintf("更新角色状态失败，失败原因：%s ", err.Error()))
		return
	}
	e.Result(req.GetId(), fmt.Sprintf("更新角色 %v 状态成功！", req.GetId()))
}

// Update2DataScope 更新角色数据权限
// @Summary 更新角色数据权限
// @Description 获取JSON
// @Tags 角色/Role
// @Accept  application/json
// @Product application/json
// @Param data body dto.RoleDataScopeReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/role-status/{id} [put]
// @Security Bearer
func (e SysRoleHandler) Update2DataScope(c *gin.Context) {
	req := dto.RoleDataScopeReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON, nil).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	data := &models.SysRole{
		RoleId:    req.RoleId,
		DataScope: req.DataScope,
		DeptIds:   req.DeptIds,
	}
	data.UpdateBy = oauth.GetUserId(c)
	err = e.biz.UpdateDataScope(&req).Error
	if err != nil {
		e.ErrorResult(500, err, fmt.Sprintf("更新角色数据权限失败！错误详情：%s", err.Error()))
		return
	}
	e.Result(nil, "操作成功")
}
