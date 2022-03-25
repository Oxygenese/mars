package task

import (
	"github.com/mars-projects/mars/api"
	"github.com/mars-projects/mars/app/system/internal/biz"
	"github.com/mars-projects/mars/app/system/internal/dto"
	"github.com/mars-projects/mars/common/transaction"
)

type QuerySysDeptTreeExecutor struct {
	*biz.SysDept
}

func (e QuerySysDeptTreeExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysDeptGetPageReq{}
	err := message.UnMarshal(&req)
	e.Log.Info("[QuerySysDeptTreeExecutor] message.UnMarshal SysDeptGetPageReq :%s", req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	list := make([]dto.DeptLabel, 0)
	list, err = e.SetDeptTree(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("查询成功", message.RequestId, &list)
	return nil
}
