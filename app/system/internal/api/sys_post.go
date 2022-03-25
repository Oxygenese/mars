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

func NewSysPostHandler(post *biz.SysPost) *SysPostHandler {
	return &SysPostHandler{biz: post, Api: api.Api{}}
}

type SysPostHandler struct {
	api.Api
	biz *biz.SysPost
}

// GetPage
// @Summary 岗位列表数据
// @Description 获取JSON
// @Tags 岗位
// @Param postName query string false "postName"
// @Param postCode query string false "postCode"
// @Param postId query string false "postId"
// @Param status query string false "status"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post [get]
// @Security Bearer
func (e SysPostHandler) GetPage(c *gin.Context) {
	req := dto.SysPostPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}

	list := make([]models.SysPost, 0)
	var count int64

	err = e.biz.GetPage(&req, &list, &count)
	if err != nil {
		e.ErrorResult(500, err, "查询失败")
		return
	}

	e.PageResult(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get
// @Summary 获取岗位信息
// @Description 获取JSON
// @Tags 岗位
// @Param id path int true "编码"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post/{postId} [get]
// @Security Bearer
func (e SysPostHandler) Get(c *gin.Context) {
	req := dto.SysPostGetReq{}
	err := e.MakeContext(c).
		Bind(&req, nil).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	var object models.SysPost

	err = e.biz.Get(&req, &object)
	if err != nil {
		e.ErrorResult(500, err, fmt.Sprintf("岗位信息获取失败！错误详情：%s", err.Error()))
		return
	}

	e.Result(object, "查询成功")
}

// Insert
// @Summary 添加岗位
// @Description 获取JSON
// @Tags 岗位
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysPostInsertReq true "data"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post [post]
// @Security Bearer
func (e SysPostHandler) Insert(c *gin.Context) {
	req := dto.SysPostInsertReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	req.SetCreateBy(c.GetInt(authentication.UserId))
	err = e.biz.Insert(&req)
	if err != nil {
		e.ErrorResult(500, err, fmt.Sprintf("新建岗位失败！错误详情：%s", err.Error()))
		return
	}
	e.Result(req.GetId(), "创建成功")
}

// Update
// @Summary 修改岗位
// @Description 获取JSON
// @Tags 岗位
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysPostUpdateReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post/{id} [put]
// @Security Bearer
func (e SysPostHandler) Update(c *gin.Context) {
	req := dto.SysPostUpdateReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON, nil).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}

	req.SetUpdateBy(c.GetInt(authentication.UserId))

	err = e.biz.Update(&req)
	if err != nil {
		e.ErrorResult(500, err, fmt.Sprintf("岗位更新失败！错误详情：%s", err.Error()))
		return
	}
	e.Result(req.GetId(), "更新成功")
}

// Delete
// @Summary 删除岗位
// @Description 删除数据
// @Tags 岗位
// @Param id body dto.SysPostDeleteReq true "请求参数"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/post [delete]
// @Security Bearer
func (e SysPostHandler) Delete(c *gin.Context) {
	req := dto.SysPostDeleteReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	req.SetUpdateBy(c.GetInt(authentication.UserId))
	err = e.biz.Remove(&req)
	if err != nil {
		e.ErrorResult(500, err, fmt.Sprintf("岗位删除失败！错误详情：%s", err.Error()))
		return
	}
	e.Result(req.GetId(), "删除成功")
}
