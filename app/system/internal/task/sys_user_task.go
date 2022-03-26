package task

import (
	"errors"
	"fmt"
	"github.com/mars-projects/mars/api"
	"github.com/mars-projects/mars/app/system/internal/biz"
	"github.com/mars-projects/mars/app/system/internal/dto"
	"github.com/mars-projects/mars/app/system/internal/models"
	"github.com/mars-projects/mars/common/transaction"
	"gorm.io/gorm"
)

type UpdateSysUserExecutor struct {
	*biz.SysUser
}

func (executor UpdateSysUserExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysUserUpdateReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	req.SetCreateBy(message.GetUserId())
	err = executor.Update(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("更新成功", message.GetRequestId(), nil)
	return nil
}

type DeleteSysUserExecutor struct {
	*biz.SysUser
}

func (executor DeleteSysUserExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysUserById{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	req.SetCreateBy(message.GetUserId())

	err = executor.Remove(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("删除成功", message.GetRequestId(), nil)
	return nil
}

type ResetSysUserPwdExecutor struct {
	*biz.SysUser
}

func (executor ResetSysUserPwdExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.ResetSysUserPwdReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	req.SetCreateBy(message.GetUserId())
	err = executor.ResetPwd(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("更新成功", message.GetRequestId(), nil)
	return nil
}

// CreateSysUserExecutor 创建用户
type CreateSysUserExecutor struct {
	*biz.SysUser
}

func (executor *CreateSysUserExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) (err error) {
	req := dto.SysUserInsertReq{}
	err = message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	// 设置创建人
	req.SetCreateBy(message.GetUserId())
	err = executor.Insert(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("创建成功", message.GetRequestId(), nil)
	return
}

type QuerySysUserProfileExecutor struct {
	*biz.SysUser
}

func (e QuerySysUserProfileExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysUserById{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	req.Id = message.GetUserId()
	sysUser := models.SysUser{}
	roles := make([]models.SysRole, 0)
	posts := make([]models.SysPost, 0)
	err = e.GetProfile(&req, &sysUser, &roles, &posts)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	res := transaction.H{
		"user":  sysUser,
		"roles": roles,
		"posts": posts,
	}
	respChan <- api.ReplyOk("个人信息查询成功", message.GetRequestId(), res)
	return nil
}

type UpdateSysUserPwdExecutor struct {
	*biz.SysUser
}

func (e UpdateSysUserPwdExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.ResetSysUserPwdReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	req.SetCreateBy(message.GetUserId())

	err = e.ResetPwd(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("密码更新成功", message.GetRequestId(), nil)
	return nil
}

type ChangeSysUserStatus struct {
	*biz.SysUser
}

func (e ChangeSysUserStatus) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.UpdateSysUserStatusReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	req.SetCreateBy(message.GetUserId())

	err = e.UpdateStatus(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("更新用户状态成功", message.GetRequestId(), nil)
	return nil
}

type QuerySysUserByIdExecutor struct {
	*biz.SysUser
}

func (e QuerySysUserByIdExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	var err error
	var data dto.SysUserById
	var model models.SysUser
	err = message.UnMarshal(&data)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	err = e.Get(&data, &model)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		respChan <- api.ReplyError(err, message.GetRequestId(), 404)
		return nil
	}
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyOk("请求成功", message.GetRequestId(), &model)
	return nil
}

// FindSysUserExecutor 查询带密码的用户信息
type FindSysUserExecutor struct {
	biz *biz.SysUser
}

func (f FindSysUserExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	var err error
	var req dto.SysUserByUsernameReq
	var model models.SysUserWithPassword
	err = message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	err = f.biz.FindSysUser(&req, &model)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	fmt.Printf("[FindSysUserExecutor] get user:%v\n", model)
	respChan <- api.ReplyOk("", message.GetRequestId(), &model)
	return nil
}

type QuerySysUserPageExecutor struct {
	*biz.SysUser
}

func (e QuerySysUserPageExecutor) Execute(message *api.Message, respChan chan *api.Reply, sender transaction.Sender) error {
	req := dto.SysUserGetPageReq{}
	err := message.UnMarshal(&req)
	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}

	list := make([]models.SysUser, 0)
	var count int64

	err = e.GetPage(&req, &list, &count)

	if err != nil {
		respChan <- api.ReplyError(err, message.GetRequestId(), 400)
		return nil
	}
	respChan <- api.ReplyPage(&list, int(count), req.PageIndex, req.PageSize, message.RequestId)
	return err
}

// SysUserInfoExecutor 登录成功后查询用户信息
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
	var mp = make(transaction.H)
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
