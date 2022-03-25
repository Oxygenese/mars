package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mars-projects/mars/app/system/internal/biz"
	"github.com/mars-projects/mars/app/system/internal/dto"
	"github.com/mars-projects/mars/app/system/internal/models"
	"github.com/mars-projects/mars/common/api"
	"github.com/mars-projects/mars/common/middleware/authentication"
)

type SysDictDataHandler struct {
	api.Api
	biz *biz.SysDictData
}

func NewSysDictDataHandler(dictData *biz.SysDictData) *SysDictDataHandler {
	return &SysDictDataHandler{biz: dictData, Api: api.Api{}}
}

// GetPage
// @Summary 字典数据列表
// @Description 获取JSON
// @Tags 字典数据
// @Param status query string false "status"
// @Param dictCode query string false "dictCode"
// @Param dictType query string false "dictType"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/data/page [get]
// @Security Bearer
func (e SysDictDataHandler) GetPage(c *gin.Context) {
	req := dto.SysDictDataGetPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}

	list := make([]models.SysDictData, 0)
	var count int64
	err = e.biz.GetPage(&req, &list, &count)
	if err != nil {
		e.ErrorResult(500, err, "查询失败")
		return
	}

	e.PageResult(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get
// @Summary 通过编码获取字典数据
// @Description 获取JSON
// @Tags 字典数据
// @Param dictCode path int true "字典编码"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/data/{dictCode} [get]
// @Security Bearer
func (e SysDictDataHandler) Get(c *gin.Context) {
	req := dto.SysDictDataGetReq{}
	err := e.MakeContext(c).
		Bind(&req, nil).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}

	var object models.SysDictData

	err = e.biz.Get(&req, &object)
	if err != nil {
		e.ErrorResult(500, err, "查询失败")
		return
	}

	e.Result(object, "查询成功")
}

// Insert
// @Summary 添加字典数据
// @Description 获取JSON
// @Tags 字典数据
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysDictDataInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/dict/data [post]
// @Security Bearer
func (e SysDictDataHandler) Insert(c *gin.Context) {
	req := dto.SysDictDataInsertReq{}
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
		e.ErrorResult(500, err, "创建失败")
		return
	}

	e.Result(req.GetId(), "创建成功")
}

// Update
// @Summary 修改字典数据
// @Description 获取JSON
// @Tags 字典数据
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysDictDataUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/dict/data/{dictCode} [put]
// @Security Bearer
func (e SysDictDataHandler) Update(c *gin.Context) {
	req := dto.SysDictDataUpdateReq{}
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
		e.ErrorResult(500, err, "更新失败")
		return
	}
	e.Result(req.GetId(), "更新成功")
}

// Delete
// @Summary 删除字典数据
// @Description 删除数据
// @Tags 字典数据
// @Param dictCode body dto.SysDictDataDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/dict/data [delete]
// @Security Bearer
func (e SysDictDataHandler) Delete(c *gin.Context) {
	req := dto.SysDictDataDeleteReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.JSON, nil).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	req.SetUpdateBy(c.GetInt(authentication.UserId))
	err = e.biz.Remove(&req)
	if err != nil {
		e.ErrorResult(500, err, "删除失败")
		return
	}
	e.Result(req.GetId(), "删除成功")
}

// GetAll 数据字典根据key获取 业务页面使用
// @Summary 数据字典根据key获取
// @Description 数据字典根据key获取
// @Tags 字典数据
// @Param dictType query int true "dictType"
// @Success 200 {object} response.Response{data=[]dto.SysDictDataGetAllResp}  "{"code": 200, "data": [...]}"
// @Router /api/v1/dict-data/option-select [get]
// @Security Bearer
func (e SysDictDataHandler) GetAll(c *gin.Context) {
	req := dto.SysDictDataGetPageReq{}
	err := e.MakeContext(c).
		Bind(&req, binding.Form).
		Errors
	if err != nil {
		e.ErrorResult(500, err, err.Error())
		return
	}
	list := make([]models.SysDictData, 0)
	err = e.biz.GetAll(&req, &list)
	if err != nil {
		e.ErrorResult(500, err, "查询失败")
		return
	}
	l := make([]dto.SysDictDataGetAllResp, 0)
	for _, i := range list {
		d := dto.SysDictDataGetAllResp{}
		e.Translate(i, &d)
		l = append(l, d)
	}

	e.Result(l, "查询成功")
}
