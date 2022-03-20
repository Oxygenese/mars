package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mars-projects/mars/common/response"
	"github.com/mars-projects/mars/common/utils"
	"net/http"
)

type Api struct {
	Context *gin.Context
	Errors  error
}

func (e *Api) AddError(err error) {
	if e.Errors == nil {
		e.Errors = err
	} else if err != nil {
		e.Errors = fmt.Errorf("%v; %w", e.Errors, err)
	}
}
func (e Api) Translate(form, to interface{}) {
	utils.Translate(form, to)
}

// MakeContext 设置http上下文
func (e *Api) MakeContext(c *gin.Context) *Api {
	e.Context = c
	return e
}

func (e *Api) Bind(d interface{}, bindings ...binding.Binding) *Api {
	var err error
	if len(bindings) == 0 {
		bindings = append(bindings, binding.JSON, nil)
	}
	for i := range bindings {
		switch bindings[i] {
		case binding.JSON:
			err = e.Context.ShouldBindWith(d, binding.JSON)
		case binding.XML:
			err = e.Context.ShouldBindWith(d, binding.XML)
		case binding.Form:
			err = e.Context.ShouldBindWith(d, binding.Form)
		case binding.Query:
			err = e.Context.ShouldBindWith(d, binding.Query)
		case binding.FormPost:
			err = e.Context.ShouldBindWith(d, binding.FormPost)
		case binding.FormMultipart:
			err = e.Context.ShouldBindWith(d, binding.FormMultipart)
		case binding.ProtoBuf:
			err = e.Context.ShouldBindWith(d, binding.ProtoBuf)
		case binding.MsgPack:
			err = e.Context.ShouldBindWith(d, binding.MsgPack)
		case binding.YAML:
			err = e.Context.ShouldBindWith(d, binding.YAML)
		case binding.Header:
			err = e.Context.ShouldBindWith(d, binding.Header)
		default:
			err = e.Context.ShouldBindUri(d)
		}
		if err != nil {
			e.AddError(err)
		}
	}
	return e
}

// ErrorResult 通常错误数据处理
func (e Api) ErrorResult(code int, err error, msg string) {
	response.Error(e.Context, code, err, msg)
}

func (e Api) Error(err error) {
	e.Context.JSON(200, err)
}

// ParamsErrorResult 通常错误数据处理
func (e Api) ParamsErrorResult(code int, err error, msg string) {
	response.Error(e.Context, code, err, msg)
}

// InternalErrorResult 通常错误数据处理
func (e Api) InternalErrorResult(err error) {
	response.Error(e.Context, http.StatusInternalServerError, err, "内部错误")
}

// Result 通常成功数据处理
func (e Api) Result(data interface{}, msg string) {
	response.OK(e.Context, data, msg)
}

// PageResult 分页数据处理
func (e Api) PageResult(result interface{}, count int, pageIndex int, pageSize int, msg string) {
	response.PageOK(e.Context, result, count, pageIndex, pageSize, msg)
}

// Custom 兼容函数
func (e Api) Custom(data gin.H) {
	response.Custum(e.Context, data)
}
