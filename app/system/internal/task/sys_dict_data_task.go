package task

import (
	"github.com/mars-projects/mars/api"
	"github.com/mars-projects/mars/app/system/internal/biz"
	"github.com/mars-projects/mars/app/system/internal/dto"
	"github.com/mars-projects/mars/app/system/internal/models"
	"github.com/mars-projects/mars/common/transaction"
	"github.com/mars-projects/mars/common/utils"
)

type SysDictDataExecutor struct {
	*biz.SysDictData
}

//QueryDictDataSelect 根据字典类型查询字典数据信息
func (executor *SysDictDataExecutor) QueryDictDataSelect(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	var err error
	req := dto.SysDictDataGetPageReq{}
	err = message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	list := make([]models.SysDictData, 0)
	err = executor.GetAll(&req, &list)
	l := make([]dto.SysDictDataGetAllResp, 0)
	for _, i := range list {
		d := dto.SysDictDataGetAllResp{}
		utils.Translate(i, &d)
		l = append(l, d)
	}
	respChan <- api.ReplyOk("查询成功", message.RequestId, &l)
	return nil
}

func (executor *SysDictDataExecutor) QueryDictDataByCode(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysDictDataGetReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	var object models.SysDictData
	err = executor.Get(&req, &object)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("查询成功", message.RequestId, &object)
	return nil
}

func (executor *SysDictDataExecutor) QueryDictDataPage(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysDictDataGetPageReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	list := make([]models.SysDictData, 0)
	var count int64
	err = executor.GetPage(&req, &list, &count)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyPage(&list, int(count), req.PageIndex, req.PageSize, message.RequestId)
	return nil
}

func (executor *SysDictDataExecutor) CreateDictData(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysDictDataInsertReq{}
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

func (executor *SysDictDataExecutor) UpdateDictData(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysDictDataUpdateReq{}
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

func (executor *SysDictDataExecutor) DeleteDictData(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysDictDataDeleteReq{}
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
	respChan <- api.ReplyOk("删除成功", message.RequestId, "")
	return nil
}
