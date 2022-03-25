package task

import (
	"github.com/mars-projects/mars/api"
	"github.com/mars-projects/mars/app/system/internal/biz"
	"github.com/mars-projects/mars/common/transaction"
)

// GetSysMenuRoleExecutor 登录成功后获取菜单路由信息
type GetSysMenuRoleExecutor struct {
	menuBiz *biz.SysMenu
}

func (g GetSysMenuRoleExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	result, err := g.menuBiz.SetMenuRole("admin")
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("请求成功", message.GetRequestId(), &result)
	return nil
}
