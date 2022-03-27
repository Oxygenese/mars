package task

import (
	"fmt"
	"github.com/mars-projects/mars/api"
	"github.com/mars-projects/mars/app/system/internal/biz"
	"github.com/mars-projects/mars/app/system/internal/dto"
	"github.com/mars-projects/mars/app/system/internal/models"
	"github.com/mars-projects/mars/common/transaction"
)

type SysRoleExecutor struct {
	*biz.SysRole
}

func (executor *SysRoleExecutor) QuerySysRolePage(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysRoleGetPageReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	list := make([]models.SysRole, 0)
	var count int64
	err = executor.GetPage(&req, &list, &count)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyPage(&list, int(count), req.PageIndex, req.PageSize, message.RequestId)
	return err
}

func (executor *SysRoleExecutor) QuerySysRoleById(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysRoleGetReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	var object models.SysRole
	err = executor.Get(&req, &object)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("查询成功", message.GetRequestId(), &object)
	return err
}

func (executor *SysRoleExecutor) CreateSysRole(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysRoleInsertReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	// 设置创建人
	req.CreateBy = message.GetUserId()
	if req.Status == "" {
		req.Status = "2"
	}
	err = executor.Insert(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("创建成功", message.GetRequestId(), nil)
	return err
}

func (executor *SysRoleExecutor) UpdateSysRole(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysRoleUpdateReq{}
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
	respChan <- api.ReplyOk("更新成功", message.GetRequestId(), "")
	return nil
}

func (executor *SysRoleExecutor) ChangeSysRoleStatus(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.UpdateStatusReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	req.SetUpdateBy(message.GetUserId())
	err = executor.UpdateStatus(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk(fmt.Sprintf("更新角色 %v 状态成功！", req.GetId()), message.GetRequestId(), nil)
	return err
}

func (executor *SysRoleExecutor) DeleteSysRole(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysRoleDeleteReq{}
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
	respChan <- api.ReplyOk(fmt.Sprintf("删除角色角色 %v 成功！", req.GetId()), message.GetRequestId(), nil)
	return err
}

func (executor *SysRoleExecutor) UpdateSysRoleDataScope(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.RoleDataScopeReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	data := &models.SysRole{
		RoleId:    req.RoleId,
		DataScope: req.DataScope,
		DeptIds:   req.DeptIds,
	}
	data.UpdateBy = message.GetUserId()
	err = executor.UpdateDataScope(&req).Error
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("操作成功", message.GetRequestId(), nil)
	return err
}
