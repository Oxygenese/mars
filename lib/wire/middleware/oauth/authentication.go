package oauth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/mars-projects/oauth2/v4"
	"net/http"
	"strconv"
	"strings"
)

const (
	UserId              = "user_id"
	UserInfo            = "user"
	TokenType           = "Bearer "
	AccessToken         = "access_token"
	AuthorizationHeader = "Authorization"

	UnAuthorized = "UnAuthorized"
)

var ProviderAuthenticationSet = wire.NewSet(NewAuthentication)

func NewAuthentication(store oauth2.TokenStore, logger log.Logger) *Authentication {
	return &Authentication{tokenStore: store, log: log.NewHelper(logger)}
}

func (e *Authentication) GinAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token, err := extractAccessToken(context)
		if err != nil {
			e.log.Debug("authentication err : %s", err)
			context.AbortWithStatusJSON(401, err)
			return
		}
		tokenInfo, err := e.loadAccessToken(token, context.Request.Context())
		if err != nil {
			e.log.Debug("authentication load access  err : %s", err)
			context.AbortWithStatusJSON(401, err)
			return
		}
		if tokenInfo == nil {
			context.AbortWithStatusJSON(401, errors.New(http.StatusUnauthorized, UnAuthorized, "请登录"))
			return
		}
		context.Set(UserId, tokenInfo.GetUserID())
		context.Set(UserInfo, tokenInfo.GetExtensionClaims())
		context.Next()
	}
}

func GetUserId(ctx *gin.Context) int {
	uid, exist := ctx.Get(UserId)
	if exist {
		op, ok := uid.(string)
		if ok {
			i, _ := strconv.Atoi(op)
			return i
		}
	}
	return 0
}

type Authentication struct {
	tokenStore oauth2.TokenStore
	log        *log.Helper
}

func (e *Authentication) loadAccessToken(access string, ctx context.Context) (oauth2.TokenInfo, error) {
	info, err := e.tokenStore.GetByAccess(ctx, access)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func extractAccessToken(ctx *gin.Context) (string, error) {
	var token string
	switch true {
	case ctx.GetHeader(AuthorizationHeader) != "":
		token = ctx.GetHeader(AuthorizationHeader)
		if strings.Contains(token, TokenType) {
			token = strings.Replace(token, TokenType, "", 1)
		}
	case ctx.Param(AccessToken) != "":
		token = ctx.Param(AccessToken)
	default:
		return "", errors.New(401, UnAuthorized, "请登录")
	}
	return token, nil
}
