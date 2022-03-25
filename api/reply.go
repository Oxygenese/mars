package api

import (
	"encoding/json"
)

func ReplyOk(message, reqId string, data interface{}) *Reply {
	res := &Reply{}
	if data != nil {
		marshal, err := json.Marshal(&data)
		if err != nil {
			res.Code = 400
			res.Message = err.Error()
			res.RequestId = reqId
		}
		res.Data = string(marshal)
	}
	res.RequestId = reqId
	res.Code = 200
	res.Message = message
	return res
}

func ReplyError(err error, requestId string, code uint32) *Reply {
	return &Reply{
		Code:      code,
		RequestId: requestId,
		Message:   err.Error(),
	}
}

type Page struct {
	PageIndex int
	PageSize  int
	Count     int
}

type page struct {
	Page
	List interface{} `json:"list"`
}

func ReplyPage(list interface{}, count int, pageIndex, pageSize int, requestId string) *Reply {
	p := &page{
		Page: Page{
			PageIndex: pageIndex,
			PageSize:  pageSize,
			Count:     count,
		},
		List: list,
	}
	return ReplyOk("查询成功", requestId, p)
}
