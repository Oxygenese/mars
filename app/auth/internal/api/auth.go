package api

import (
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/mars-projects/mars/common/transaction"
	"github.com/mars-projects/oauth2/v4"
)

func (api *TokenApi) Token(ctx http.Context) error {
	gt, tgr, err := api.server.ValidationTokenRequest(ctx.Request())
	if err != nil {
		log.Errorf("[auth] ValidationTokenRequest err:%s", err)
		return err
	}
	ti, err := api.server.GetAccessToken(ctx, gt, tgr)
	if err != nil {
		log.Errorf("[auth] GetAccessToken err:%s", err)
		return err
	}
	data := api.server.GetTokenData(ti)
	d, err := json.Marshal(&data)
	if err != nil {
		log.Errorf("[auth] marshal token err:%s", data)
		return err
	}
	var res = transaction.H{
		"code": 200,
		"data": d,
		"msg":  "认证成功",
	}
	return ctx.JSON(200, &res)
}

func (api *TokenApi) Logout(ctx http.Context) error {
	accessToken := ctx.Header().Get("Authorization")
	if accessToken == "" {
		accessToken = ctx.Query().Get("token")
	}
	err := api.server.Manager.RemoveAccessToken(ctx, accessToken)
	if err != nil {
		return err
	}
	var res = transaction.H{
		"code":    200,
		"message": "注销成功",
		"data":    "",
	}
	return ctx.JSON(200, res)
}

func (api *TokenApi) Authorize(ctx http.Context) error {
	err := api.server.HandleAuthorizeRequest(ctx.Response(), ctx.Request())
	if err != nil {
		return err
	}
	return nil
}

func (api *TokenApi) Confirm(ctx http.Context) error {
	accessToken := ctx.Header().Get("Authorization")
	if accessToken == "" {
		accessToken = ctx.Query().Get("token")
	}
	data, err := api.server.Manager.LoadAccessToken(ctx, accessToken)
	if err != nil {
		return err
	}
	var res = map[string]interface{}{
		"code":    200,
		"data":    data,
		"message": "登录成功",
	}
	return ctx.JSON(200, res)
}

func (api *TokenApi) RefreshToken(ctx http.Context) error {
	refreshToken := ctx.Query().Get("refresh_token")
	res, err := api.server.Manager.LoadRefreshToken(ctx, refreshToken)
	if err != nil {
		return err
	}
	ti, err := api.server.GetAccessToken(ctx, "refresh_token", &oauth2.TokenGenerateRequest{
		ClientID:            res.GetClientID(),
		UserID:              res.GetUserID(),
		RedirectURI:         res.GetRedirectURI(),
		Scope:               res.GetScope(),
		Code:                res.GetCode(),
		CodeChallenge:       res.GetCodeChallenge(),
		CodeChallengeMethod: res.GetCodeChallengeMethod(),
		Refresh:             res.GetRefresh(),
		AccessTokenExp:      res.GetAccessExpiresIn(),
		Request:             ctx.Request(),
	})
	if err != nil {
		return err
	}
	return ctx.JSON(200, api.server.GetTokenData(ti))
}
