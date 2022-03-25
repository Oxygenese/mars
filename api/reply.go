package api

import (
	"encoding/json"
)

func ReplyOk(message, reqId string, data interface{}) *Reply {
	res := &Reply{}
	marshal, err := json.Marshal(&data)
	if err != nil {
		res.Code = 400
		res.Message = err.Error()
		res.RequestId = reqId
	}
	res.RequestId = reqId
	res.Code = 200
	res.Message = message
	res.Data = string(marshal)
	return res
}

func ReplyError(err error, requestId string, code uint32) *Reply {
	return &Reply{
		Code:      code,
		RequestId: requestId,
		Message:   err.Error(),
	}
}
