package task

import (
	"github.com/mars-projects/mars/api"
	"github.com/mars-projects/mars/app/system/internal/biz"
	"github.com/mars-projects/mars/app/system/internal/dto"
	"github.com/mars-projects/mars/app/system/internal/models"
	"github.com/mars-projects/mars/common/transaction"
)

type SysDictTypeExecutor struct {
	*biz.SysDictType
}

func (executor *SysDictTypeExecutor) QueryDictTypePage(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysDictTypeGetPageReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	list := make([]models.SysDictType, 0)
	var count int64
	err = executor.GetPage(&req, &list, &count)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyPage(&list, int(count), req.PageIndex, req.PageSize, message.RequestId)
	return nil
}

func (executor *SysDictTypeExecutor) QueryDictTypeById(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysDictTypeGetReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	var object models.SysDictType
	err = executor.Get(&req, &object)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("查询成功", message.RequestId, &object)
	return nil
}

func (executor *SysDictTypeExecutor) CreateDictType(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysDictTypeInsertReq{}
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

func (executor *SysDictTypeExecutor) UpdateDictType(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysDictTypeUpdateReq{}
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
	respChan <- api.ReplyOk("更新成功", message.RequestId, nil)
	return nil
}

func (executor *SysDictTypeExecutor) DeleteDictType(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysDictTypeDeleteReq{}
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

func (executor *SysDictTypeExecutor) ExportDictType(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	//TODO implement me
	panic("implement me")
}

func (executor *SysDictTypeExecutor) QueryDictTypeOptionSelect(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysDictTypeGetPageReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	list := make([]models.SysDictType, 0)
	err = executor.GetAll(&req, &list)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("查询成功", message.RequestId, nil)
	return nil
}
