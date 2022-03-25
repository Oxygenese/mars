package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mars-projects/mars/api"
)

type UserBiz struct {
	client api.SystemClient
	log    *log.Helper
}

func NewUserBiz(client api.SystemClient, logger log.Logger) *UserBiz {
	return &UserBiz{client: client, log: log.NewHelper(log.With(logger, "modules", "usecase/sys"))}
}

func (uc UserBiz) FindSysUser(ctx context.Context, in *api.Request) (*api.Reply, error) {
	return uc.client.OnMessageReceived(ctx, in)
}
