package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mars-projects/mars/api/system"
)

type UserBiz struct {
	client system.SystemClient
	log    *log.Helper
}

func NewUserBiz(client system.SystemClient, logger log.Logger) *UserBiz {
	return &UserBiz{client: client, log: log.NewHelper(log.With(logger, "module", "usecase/sys"))}
}

func (uc UserBiz) FindSysUser(ctx context.Context, in *system.SysUserInfoReq) (*system.SysUserReply, error) {
	return uc.client.SysUserInfo(ctx, in)
}
