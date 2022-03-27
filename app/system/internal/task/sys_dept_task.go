package task

import (
	"github.com/mars-projects/mars/api"
	"github.com/mars-projects/mars/app/system/internal/biz"
	"github.com/mars-projects/mars/app/system/internal/dto"
	"github.com/mars-projects/mars/app/system/internal/models"
	"github.com/mars-projects/mars/common/transaction"
)

type SysDeptExecutor struct {
	*biz.SysDept
}

func (executor *SysDeptExecutor) DeleteSysDept(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysDeptDeleteReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	err = executor.Remove(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("删除成功", message.RequestId, nil)
	return nil
}

func (executor *SysDeptExecutor) UpdateSysDept(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysDeptUpdateReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	req.SetUpdateBy(message.GetUserId())
	err = executor.Update(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("更新成功", message.RequestId, "")
	return nil
}

func (executor *SysDeptExecutor) QuerySysDeptPage(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysDeptGetPageReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	list := make([]models.SysDept, 0)
	list, err = executor.SetDeptPage(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("查询成功", message.RequestId, &list)
	return nil
}

func (executor *SysDeptExecutor) QuerySysDeptById(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysDeptGetReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	var object models.SysDept
	err = executor.Get(&req, &object)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("查询成功", message.RequestId, &object)
	return nil
}

func (executor *SysDeptExecutor) QuerySysDeptTree(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysDeptGetPageReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	list := make([]dto.DeptLabel, 0)
	list, err = executor.SetDeptTree(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("查询成功", message.RequestId, &list)
	return nil
}

func (executor *SysDeptExecutor) QuerySysDeptTreeRoleSelect(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SelectRole{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	result, err := executor.SetDeptLabel()
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	menuIds := make([]int, 0)
	if req.RoleId != 0 {
		menuIds, err = executor.GetWithRoleId(req.RoleId)
		if err != nil {
			respChan <- api.ReplyError(err, message.GetRequestId(), 400)
			return nil
		}
	}
	res := transaction.H{
		"depts":       result,
		"checkedKeys": menuIds,
	}
	respChan <- api.ReplyOk("查询成功", message.RequestId, &res)
	return nil
}
