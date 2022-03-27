package task

import (
	"github.com/mars-projects/mars/api"
	"github.com/mars-projects/mars/app/system/internal/biz"
	"github.com/mars-projects/mars/app/system/internal/dto"
	"github.com/mars-projects/mars/app/system/internal/models"
	"github.com/mars-projects/mars/common/transaction"
)

type SysPostExecutor struct {
	*biz.SysPost
}

func (executor *SysPostExecutor) DeleteSysPost(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysPostDeleteReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	req.SetUpdateBy(message.GetUserId())
	err = executor.Remove(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("删除成功", message.RequestId, nil)
	return nil
}

func (executor *SysPostExecutor) UpdateSysPost(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysPostUpdateReq{}
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

func (executor *SysPostExecutor) CreateSysPost(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysPostInsertReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	req.SetCreateBy(message.GetUserId())
	err = executor.Insert(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("创建成功", message.RequestId, nil)
	return nil
}

func (executor *SysPostExecutor) QuerySysPostById(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysPostGetReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	var object models.SysPost
	err = executor.Get(&req, &object)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("查询成功", message.RequestId, &object)
	return nil
}

func (executor *SysPostExecutor) QuerySysPostPage(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysPostPageReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	list := make([]models.SysPost, 0)
	var count int64

	err = executor.GetPage(&req, &list, &count)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyPage(&list, int(count), req.PageIndex, req.PageSize, message.RequestId)
	return err
}
