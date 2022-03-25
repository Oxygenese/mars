package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mars-projects/mars/app/system/internal/biz"
	"github.com/mars-projects/mars/app/system/internal/dto"
	"github.com/mars-projects/mars/app/system/internal/models"
	"github.com/mars-projects/mars/common/api"
	"github.com/mars-projects/mars/common/middleware/authentication"
	"github.com/mars-projects/mars/common/utils"
)

type SysDeptHandler struct {
	api.Api
	biz *biz.SysDept
}

func NewSysDeptHandler(dept *biz.SysDept) *SysDeptHandler {
	return &SysDeptHandler{biz: dept, Api: api.Api{}}
}

// GetPage
// @Summary 分页部门列表数据
// @Description 分页列表
// @Tags 部门
// @Param deptName query string false "deptName"
// @Param deptId query string false "deptId"
// @Param position query string false "position"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dept [get]
// @Security Bearer
func (e SysDeptHandler) GetPage(c *gin.Context) {
	req := dto.SysDeptGetPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	list := make([]models.SysDept, 0)
	list, err = e.biz.SetDeptPage(&req)
	if err != nil {
		e.ErrorResult(500, err, "查询失败")
		return
	}
	e.Result(list, "查询成功")
}

// Get
// @Summary 获取部门数据
// @Description 获取JSON
// @Tags 部门
// @Param deptId path string false "deptId"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dept/{deptId} [get]
// @Security Bearer
func (e SysDeptHandler) Get(c *gin.Context) {
	req := dto.SysDeptGetReq{}
	err := e.MakeContext(c).
		Bind(&req, nil).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}

	var object models.SysDept

	err = e.biz.Get(&req, &object)
	if err != nil {
		e.ErrorResult(500, err, "查询失败")
		return
	}

	e.Result(object, "查询成功")
}

// Insert 添加部门
// @Summary 添加部门
// @Description 获取JSON
// @Tags 部门
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysDeptInsertReq true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dept [post]
// @Security Bearer
func (e SysDeptHandler) Insert(c *gin.Context) {
	req := dto.SysDeptInsertReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	req.SetCreateBy(c.GetInt(authentication.UserId))
	// 设置创建人
	err = e.biz.Insert(&req)
	if err != nil {
		e.ErrorResult(500, err, "创建失败")
		return
	}
	e.Result(req.GetId(), "创建成功")
}

// Update
// @Summary 修改部门
// @Description 获取JSON
// @Tags 部门
// @Accept  application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SysDeptUpdateReq true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dept/{deptId} [put]
// @Security Bearer
func (e SysDeptHandler) Update(c *gin.Context) {
	req := dto.SysDeptUpdateReq{}
	err := e.MakeContext(c).
		Bind(&req).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	req.SetUpdateBy(c.GetInt(authentication.UserId))
	err = e.biz.Update(&req)
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	e.Result(req.GetId(), "更新成功")
}

// Delete
// @Summary 删除部门
// @Description 删除数据
// @Tags 部门
// @Param data body dto.SysDeptDeleteReq true "body"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/dept [delete]
// @Security Bearer
func (e SysDeptHandler) Delete(c *gin.Context) {
	req := dto.SysDeptDeleteReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON, nil).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}

	err = e.biz.Remove(&req)
	if err != nil {
		e.ErrorResult(500, err, "删除失败")
		return
	}
	e.Result(req.GetId(), "删除成功")
}

// Get2Tree 用户管理 左侧部门树
func (e SysDeptHandler) Get2Tree(c *gin.Context) {
	req := dto.SysDeptGetPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	list := make([]dto.DeptLabel, 0)
	list, err = e.biz.SetDeptTree(&req)
	if err != nil {
		e.ErrorResult(500, err, "查询失败")
		return
	}
	e.Result(list, "")
}

// GetDeptTreeRoleSelect TODO: 此接口需要调整不应该将list和选中放在一起
func (e SysDeptHandler) GetDeptTreeRoleSelect(c *gin.Context) {
	req := dto.SelectRole{}
	err := e.MakeContext(c).
		Bind(req, nil).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	id, err := utils.StringToInt(c.Param("roleId"))
	result, err := e.biz.SetDeptLabel()
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	menuIds := make([]int, 0)
	if id != 0 {
		menuIds, err = e.biz.GetWithRoleId(id)
		if err != nil {
			e.ErrorResult(500, err, err.Error())
			return
		}
	}
	e.Result(gin.H{
		"depts":       result,
		"checkedKeys": menuIds,
	}, "")
}
