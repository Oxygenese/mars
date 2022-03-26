package oauth2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/mars-projects/mars/api"
	"github.com/mars-projects/mars/app/auth/internal/biz"
	"github.com/mars-projects/oauth2/v4"
	"github.com/mars-projects/oauth2/v4/manage"
	"github.com/mars-projects/oauth2/v4/server"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

var ProviderOauth = wire.NewSet(NewManage, NewServer)

type SysUserWithPassword struct {
	UserId   int64  `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password" `
	NickName string `json:"nickName" `
	Phone    string `json:"phone"`
	RoleId   int    `json:"roleId"`
	Avatar   string `json:"avatar" `
	Sex      string `json:"sex" `
	Email    string `json:"email" `
	DeptId   int    `json:"deptId"`
	PostId   int    `json:"postId" `
	Remark   string `json:"remark"`
	Status   string `json:"status" `
}

func NewServer(manager *manage.Manager, biz *biz.UserBiz) *server.Server {
	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	srv.SetPasswordAuthorizationHandler(
		func(username, password string) (userID string, err error) {
			user := SysUserWithPassword{Username: username}
			marshal, err := json.Marshal(&user)
			if err != nil {
				log.Errorf("[oauth] Marshal SysUserWithPassword err :%s", err)
				return "", err
			}
			req := &api.Request{
				Data:    marshal,
				Operate: api.Operate_FindSysUser,
			}
			res, err := biz.FindSysUser(context.Background(), req)
			if err != nil {
				log.Errorf("[oauth] FindSysUser failed :%s", err)
				return "", errors.New(400, "BadRequest", err.Error())
			}
			err = json.Unmarshal(res.Data, &user)
			fmt.Printf("[oauth] get user :%v\n", user)
			if err != nil {
				log.Errorf("[oauth] Unmarshal SysUserWithPassword err :%s", err)
				return "", err
			}
			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if err != nil {
				log.Info(err)
				return "", errors.New(401, "Incorrect Username or Password", "用户名或密码错误")
			}
			srv.SetExtensionClaimHandler(
				func(tgr *oauth2.TokenGenerateRequest) {
					tgr.ExtensionClaims = user
				},
			)
			return strconv.FormatInt(user.UserId, 10), nil
		},
	)
	return srv
}
