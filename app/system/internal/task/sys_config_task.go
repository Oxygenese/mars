package task

import (
	"github.com/mars-projects/mars/api"
	"github.com/mars-projects/mars/app/system/internal/biz"
	"github.com/mars-projects/mars/app/system/internal/dto"
	"github.com/mars-projects/mars/app/system/internal/models"
	"github.com/mars-projects/mars/common/transaction"
)

// SysConfigExecutor 查询系统设置分页数据
type SysConfigExecutor struct {
	*biz.SysConfig
}

func (executor *SysConfigExecutor) QuerySysConfigPage(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	var err error
	req := dto.SysConfigGetPageReq{}
	err = message.UnMarshal(&req)
	executor.Log.Debug("[QuerySysConfigPageExecutor] unmarshal SysConfigGetPageReq :%s", req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	list := make([]models.SysConfig, 0)
	var count int64
	err = executor.GetPage(&req, &list, &count)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyPage(&list, int(count), req.PageIndex, req.GetPageIndex(), message.RequestId)
	return nil
}

func (executor *SysConfigExecutor) CreateSysConfig(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	var err error
	req := dto.SysConfigControl{}
	err = message.UnMarshal(&req)
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
	respChan <- api.ReplyOk("创建成功", message.GetRequestId(), nil)
	return err
}

func (executor *SysConfigExecutor) QueryAppConfig(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) (err error) {
	req := dto.SysConfigGetToSysAppReq{}
	err = message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	req.IsFrontend = 1
	list := make([]models.SysConfig, 0)
	err = executor.GetWithKeyList(&req, &list)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	mp := make(transaction.H)
	for i := 0; i < len(list); i++ {
		key := list[i].ConfigKey
		if key != "" {
			mp[key] = list[i].ConfigValue
		}
	}
	respChan <- api.ReplyOk("请求成功", message.GetRequestId(), &mp)
	return nil
}

func (executor *SysConfigExecutor) QuerySysConfigSet(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := make([]dto.GetSetSysConfigReq, 0)
	err := executor.GetForSet(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	m := make(transaction.H, 0)
	for _, v := range req {
		m[v.ConfigKey] = v.ConfigValue
	}
	respChan <- api.ReplyOk("请求成功", message.GetRequestId(), &m)
	return nil
}

func (executor *SysConfigExecutor) UpdateSysConfigSet(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	var err error
	req := make([]dto.GetSetSysConfigReq, 0)
	err = message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	err = executor.UpdateForSet(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("系统设置更新成功", message.GetRequestId(), nil)
	return nil
}

func (executor *SysConfigExecutor) QuerySysConfigById(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysConfigGetReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	var object models.SysConfig
	err = executor.Get(&req, &object)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("查询成功", message.GetRequestId(), &object)
	return nil
}

func (executor *SysConfigExecutor) QuerySysConfigByKey(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	var err error
	var req = new(dto.SysConfigByKeyReq)
	var resp = new(dto.GetSysConfigByKEYForServiceResp)
	err = message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	err = executor.GetWithKey(req, resp)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("查询成功", message.GetRequestId(), resp)
	return nil
}

func (executor *SysConfigExecutor) UpdateSysConfig(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysConfigControl{}
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
	respChan <- api.ReplyOk("更新成功", message.GetRequestId(), nil)
	return nil
}

func (executor *SysConfigExecutor) DeleteSysConfig(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysConfigDeleteReq{}
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
	respChan <- api.ReplyOk("删除成功", message.GetRequestId(), nil)
	return nil
}
