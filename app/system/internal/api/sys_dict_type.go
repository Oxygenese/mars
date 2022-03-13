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
)

type SysDictTypeHandler struct {
	api.Api
	biz *biz.SysDictType
}

func NewSysDictTypHandler(dictType *biz.SysDictType) *SysDictTypeHandler {
	return &SysDictTypeHandler{biz: dictType, Api: api.Api{}}
}

// GetPage 字典类型列表数据
// @Summary 字典类型列表数据
// @Description 获取JSON
// @Tags 字典类型
// @Param dictName query string false "dictName"
// @Param dictId query string false "dictId"
// @Param dictType query string false "dictType"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/type [get]
// @Security Bearer
func (e SysDictTypeHandler) GetPage(c *gin.Context) {
	req := dto.SysDictTypeGetPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	list := make([]models.SysDictType, 0)
	var count int64
	err = e.biz.GetPage(&req, &list, &count)
	if err != nil {
		e.ErrorResult(500, err, "查询失败")
		return
	}
	e.PageResult(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 字典类型通过字典id获取
// @Summary 字典类型通过字典id获取
// @Description 获取JSON
// @Tags 字典类型
// @Param dictId path int true "字典类型编码"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/type/{dictId} [get]
// @Security Bearer
func (e SysDictTypeHandler) Get(c *gin.Context) {
	req := dto.SysDictTypeGetReq{}
	err := e.MakeContext(c).
		Bind(&req, nil).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	var object models.SysDictType
	err = e.biz.Get(&req, &object)
	if err != nil {
		e.ErrorResult(500, err, "查询失败")
		return
	}
	e.Result(object, "查询成功")
}

//Insert 字典类型创建
// @Summary 添加字典类型
// @Description 获取JSON
// @Tags 字典类型
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysDictTypeInsertReq true "data"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/type [post]
// @Security Bearer
func (e SysDictTypeHandler) Insert(c *gin.Context) {
	req := dto.SysDictTypeInsertReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	req.SetCreateBy(oauth.GetUserId(c))
	err = e.biz.Insert(&req)
	if err != nil {
		e.ErrorResult(500, err, fmt.Sprintf(" 创建字典类型失败，详情：%s", err.Error()))
		return
	}
	e.Result(req.GetId(), "创建成功")
}

// Update
// @Summary 修改字典类型
// @Description 获取JSON
// @Tags 字典类型
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysDictTypeUpdateReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/type/{dictId} [put]
// @Security Bearer
func (e SysDictTypeHandler) Update(c *gin.Context) {
	req := dto.SysDictTypeUpdateReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON, nil).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	req.SetUpdateBy(oauth.GetUserId(c))
	err = e.biz.Update(&req)
	if err != nil {
		return
	}
	e.Result(req.GetId(), "更新成功")
}

// Delete
// @Summary 删除字典类型
// @Description 删除数据
// @Tags 字典类型
// @Param dictCode body dto.SysDictTypeDeleteReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/type [delete]
// @Security Bearer
func (e SysDictTypeHandler) Delete(c *gin.Context) {
	req := dto.SysDictTypeDeleteReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON, nil).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	req.SetUpdateBy(oauth.GetUserId(c))
	err = e.biz.Remove(&req)
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	e.Result(req.GetId(), "删除成功")
}

// GetAll
// @Summary 字典类型全部数据 代码生成使用接口
// @Description 获取JSON
// @Tags 字典类型
// @Param dictName query string false "dictName"
// @Param dictId query string false "dictId"
// @Param dictType query string false "dictType"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/type-option-select [get]
// @Security Bearer
func (e SysDictTypeHandler) GetAll(c *gin.Context) {
	req := dto.SysDictTypeGetPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	list := make([]models.SysDictType, 0)
	err = e.biz.GetAll(&req, &list)
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	e.Result(list, "查询成功")
}
