package task

import (
	"fmt"
	"github.com/mars-projects/mars/api"
	"github.com/mars-projects/mars/app/system/internal/biz"
	"github.com/mars-projects/mars/app/system/internal/dto"
	"github.com/mars-projects/mars/app/system/internal/models"
	"github.com/mars-projects/mars/common/transaction"
)

// QuerySysConfigPageExecutor 查询系统设置分页数据
type QuerySysConfigPageExecutor struct {
	*biz.SysConfig
}

func (e QuerySysConfigPageExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	var err error
	req := dto.SysConfigGetPageReq{}
	err = message.UnMarshal(&req)
	e.Log.Debug("[QuerySysConfigPageExecutor] unmarshal SysConfigGetPageReq :%s", req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	list := make([]models.SysConfig, 0)
	var count int64
	err = e.GetPage(&req, &list, &count)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyPage(&list, int(count), req.PageIndex, req.GetPageIndex(), message.RequestId)
	return nil
}

// CreateSysConfigExecutor 创建配置项
type CreateSysConfigExecutor struct {
	*biz.SysConfig
}

func (e CreateSysConfigExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	var err error
	req := dto.SysConfigControl{}
	err = message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	req.SetCreateBy(message.GetUserId())

	err = e.Insert(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("创建成功", message.GetRequestId(), nil)
	return err
}

type QueryAppConfigExecutor struct {
	biz *biz.SysConfig
}

func (e QueryAppConfigExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) (err error) {
	req := dto.SysConfigGetToSysAppReq{}
	fmt.Println("传入的data：", message.Data)
	err = message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	fmt.Println("请求参数", req)
	req.IsFrontend = 1
	list := make([]models.SysConfig, 0)
	err = e.biz.GetWithKeyList(&req, &list)
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
	m := make(transaction.H, 0)
	for _, v := range req {
		m[v.ConfigKey] = v.ConfigValue
	}
	respChan <- api.ReplyOk("请求成功", message.GetRequestId(), &m)
	return nil
}

// UpdateSysConfigSetExecutor 更新系统设置
type UpdateSysConfigSetExecutor struct {
	*biz.SysConfig
}

func (e UpdateSysConfigSetExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	var err error
	req := make([]dto.GetSetSysConfigReq, 0)
	err = message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	err = e.UpdateForSet(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("系统设置更新成功", message.GetRequestId(), nil)
	return nil
}

//QuerySysConfigByIdExecutor 根据id查询配置
type QuerySysConfigByIdExecutor struct {
	*biz.SysConfig
}

func (e QuerySysConfigByIdExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysConfigGetReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	var object models.SysConfig
	err = e.Get(&req, &object)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("查询成功", message.GetRequestId(), &object)
	return nil
}

// QuerySysConfigByKeyExecutor 根据字典Key获取字典数据
type QuerySysConfigByKeyExecutor struct {
	*biz.SysConfig
}

func (e QuerySysConfigByKeyExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	var err error
	var req = new(dto.SysConfigByKeyReq)
	var resp = new(dto.GetSysConfigByKEYForServiceResp)
	err = message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	err = e.GetWithKey(req, resp)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("查询成功", message.GetRequestId(), resp)
	return nil
}

// UpdateSysConfigExecutor 更新系统设置
type UpdateSysConfigExecutor struct {
	*biz.SysConfig
}

func (e UpdateSysConfigExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysConfigControl{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	req.SetUpdateBy(message.GetUserId())
	err = e.Update(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("更新成功", message.GetRequestId(), nil)
	return nil
}

type DeleteSysConfigExecutor struct {
	*biz.SysConfig
}

func (e DeleteSysConfigExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysConfigDeleteReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	req.SetUpdateBy(message.GetUserId())
	err = e.Remove(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("删除成功", message.GetRequestId(), nil)
	return nil
}
