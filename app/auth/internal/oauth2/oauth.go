package oauth2

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/mars-projects/mars/api/system"
	"github.com/mars-projects/mars/app/auth/internal/biz"
	"github.com/mars-projects/oauth2/v4"
	"github.com/mars-projects/oauth2/v4/manage"
	"github.com/mars-projects/oauth2/v4/server"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

var ProviderOauth = wire.NewSet(NewManage, NewServer)

func NewServer(manager *manage.Manager, biz *biz.UserBiz) *server.Server {
	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	srv.SetPasswordAuthorizationHandler(
		func(username, password string) (userID string, err error) {
			userInfo, err := biz.FindSysUser(context.Background(), &system.SysUserInfoReq{Username: username})
			if err != nil {
				return "", errors.New(401, err.Error(), "请求失败")
			}
			fmt.Println(userInfo)
			err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(password))
			if err != nil {
				log.Info(err)
				return "", errors.New(401, "Incorrect Username or Password", "用户名或密码错误")
			}
			srv.SetExtensionClaimHandler(
				func(tgr *oauth2.TokenGenerateRequest) {
					tgr.ExtensionClaims = userInfo
				},
			)
			return strconv.FormatInt(userInfo.UserId, 10), nil
		},
	)
	return srv
}
