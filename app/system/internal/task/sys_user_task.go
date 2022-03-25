package task

import (
	"github.com/mars-projects/mars/api"
	"github.com/mars-projects/mars/app/system/internal/biz"
	"github.com/mars-projects/mars/app/system/internal/dto"
	"github.com/mars-projects/mars/app/system/internal/models"
	"github.com/mars-projects/mars/common/transaction"
)

type CreateSysUserExecutor struct {
}

func (executor *CreateSysUserExecutor) Execute(request *api.Message, respChan chan *api.Reply, sender transaction.Sender) (err error) {
	res := &api.Reply{
		Code:    200,
		Data:    "",
		Message: "创建成功",
	}
	respChan <- res
	return nil
}

type FindSysUserExecutor struct {
	biz *biz.SysUser
}

func (f FindSysUserExecutor) Execute(msg *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	var err error
	var req dto.SysUserByUsernameReq
	var model models.SysUserWithPassword
	err = msg.UnMarshal(&req)
	err = f.biz.FindSysUser(&req, &model)
	if err != nil {
		respChan <- api.ReplyError(err, msg.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("请求成功", msg.GetRequestId(), &model)
	return nil
}

type SysUserInfoExecutor struct {
	biz *biz.SysUser
}

func (s SysUserInfoExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) (err error) {
	body := dto.SysUserById{}
	body.Id = message.GetUserId()
	var roles = make([]string, 1)
	var permissions = make([]string, 1)
	permissions[0] = "*:*:*"
	var buttons = make([]string, 1)
	buttons[0] = "*:*:*"
	roles[0] = "admin"
	var mp = make(map[string]interface{})
	mp["roles"] = roles
	mp["permissions"] = permissions
	mp["buttons"] = buttons
	sysUser := models.SysUser{}
	err = s.biz.Get(&body, &sysUser)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}

	mp["introduction"] = " am a super administrator"
	mp["avatar"] = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	if sysUser.Avatar != "" {
		mp["avatar"] = sysUser.Avatar
	}

	mp["userName"] = sysUser.NickName
	mp["userId"] = sysUser.UserId
	mp["deptId"] = sysUser.DeptId
	mp["name"] = sysUser.NickName
	mp["code"] = 200

	respChan <- api.ReplyOk("请求成功", message.GetRequestId(), &mp)
	return
}
