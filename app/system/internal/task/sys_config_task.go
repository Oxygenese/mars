package task

import (
	"github.com/mars-projects/mars/api"
	"github.com/mars-projects/mars/app/system/internal/biz"
	"github.com/mars-projects/mars/app/system/internal/dto"
	"github.com/mars-projects/mars/app/system/internal/models"
	"github.com/mars-projects/mars/common/transaction"
)

type CreateSysConfigExecutor struct {
}

func (c CreateSysConfigExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	//TODO implement me
	panic("implement me")
}

type QuerySysConfigExecutor struct {
	biz *biz.SysConfig
}

func (e QuerySysConfigExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) (err error) {
	var req dto.SysConfigGetReq
	var model models.SysConfig
	err = message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	err = e.biz.Get(&req, &model)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("请求成功", message.GetRequestId(), &model)
	return nil
}

type QueryAppConfigExecutor struct {
	biz *biz.SysConfig
}

func (e QueryAppConfigExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) (err error) {
	req := dto.SysConfigGetToSysAppReq{}
	req.IsFrontend = 1
	list := make([]models.SysConfig, 0)
	err = e.biz.GetWithKeyList(&req, &list)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	mp := make(map[string]string)
	for i := 0; i < len(list); i++ {
		key := list[i].ConfigKey
		if key != "" {
			mp[key] = list[i].ConfigValue
		}
	}
	respChan <- api.ReplyOk("请求成功", message.GetRequestId(), &mp)
	return nil
}

type QuerySysConfigSetExecutor struct {
	biz *biz.SysConfig
}

func (g QuerySysConfigSetExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := make([]dto.GetSetSysConfigReq, 0)
	err := g.biz.GetForSet(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	m := make(map[string]interface{}, 0)
	for _, v := range req {
		m[v.ConfigKey] = v.ConfigValue
	}
	respChan <- api.ReplyOk("请求成功", message.GetRequestId(), &m)
	return nil
}
