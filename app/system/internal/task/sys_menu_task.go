package task

import (
	"github.com/mars-projects/mars/api"
	"github.com/mars-projects/mars/app/system/internal/biz"
	"github.com/mars-projects/mars/app/system/internal/dto"
	"github.com/mars-projects/mars/app/system/internal/models"
	"github.com/mars-projects/mars/common/transaction"
)

type QuerySysMenuTreeSelectExecutor struct {
	roleBiz *biz.SysRole
	menuBiz *biz.SysMenu
}

func (executor QuerySysMenuTreeSelectExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SelectRole{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	result, err := executor.menuBiz.SetLabel()
	menuIds := make([]int, 0)
	if req.RoleId != 0 {
		menuIds, err = executor.roleBiz.GetRoleMenuId(req.RoleId)
		if err != nil {
			respChan <- api.ReplyError(err, message.GetRequestId(), 400)
			return nil
		}
	}
	res := map[string]interface{}{
		"menus":       result,
		"checkedKeys": menuIds,
	}
	respChan <- api.ReplyOk("查询成功", message.GetRequestId(), &res)
	return err
}

// QuerySysMenuRoleExecutor 登录成功后获取菜单路由信息
type QuerySysMenuRoleExecutor struct {
	*biz.SysMenu
}

func (executor QuerySysMenuRoleExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	result, err := executor.SetMenuRole("admin")
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("请求成功", message.GetRequestId(), &result)
	return nil
}

type QuerySysMenuPageExecutor struct {
	*biz.SysMenu
}

func (executor QuerySysMenuPageExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysMenuGetPageReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	var list = make([]models.SysMenu, 0)
	err = executor.GetPage(&req, &list).Error
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("请求成功", message.GetRequestId(), &list)
	return nil
}

type QuerySysMenuByIdExecutor struct {
	*biz.SysMenu
}

func (executor QuerySysMenuByIdExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysMenuGetReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}

	var object = models.SysMenu{}

	err = executor.Get(&req, &object).Error
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("查询成功", message.GetRequestId(), &object)
	return nil
}

type CreateSysMenuExecutor struct {
	*biz.SysMenu
}

func (executor CreateSysMenuExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysMenuInsertReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	// 设置创建人
	req.SetCreateBy(message.GetUserId())
	err = executor.Insert(&req).Error
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("创建成功", message.GetRequestId(), nil)
	return nil
}

type UpdateSysMenuExecutor struct {
	*biz.SysMenu
}

func (executor UpdateSysMenuExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysMenuUpdateReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	req.SetCreateBy(message.GetUserId())
	err = executor.Update(&req).Error
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("更新成功", message.GetRequestId(), nil)
	return nil
}

type DeleteSysMenuExecutor struct {
	*biz.SysMenu
}

func (executor DeleteSysMenuExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := new(dto.SysMenuDeleteReq)
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	req.SetUpdateBy(message.GetUserId())
	err = executor.Remove(req).Error
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("删除成功", message.GetRequestId(), nil)
	return nil
}
