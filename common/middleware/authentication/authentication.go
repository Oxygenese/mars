package authentication

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/mars-projects/mars/api"
	"github.com/mars-projects/oauth2/v4"
	"net/http"
	"strings"
)

const (
	UserId              = "user_id"
	TokenType           = "Bearer "
	AuthorizationHeader = "Authorization"
	UnAuthorized        = "UnAuthorized"
)

func logWithCtx(ctx context.Context, logger log.Logger, err error) {
	_ = log.WithContext(ctx, logger).Log(
		log.LevelError,
		"reason:",
		err,
	)
}

func CheckWhiteList(operate api.Operate) bool {
	if operate == 0 {
		return true
	}
	return false
}

func Server(store oauth2.TokenStore, logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			switch req.(type) {
			case *api.Request:
				msg := req.(*api.Request)
				if CheckWhiteList(msg.Operate) {
					return handler(ctx, msg)
				}
				tr, ok := transport.FromServerContext(ctx)
				if !ok {
					log.Errorf("[authentication] context transport get err")
					return nil, errors.New(http.StatusBadRequest, "BadRequest", "参数错误")
				}
				header := tr.RequestHeader()
				token := header.Get(AuthorizationHeader)
				if !strings.Contains(token, TokenType) {
					err := errors.New(6401, "Invalid Token Type", "未登录或登录已过期")
					logWithCtx(ctx, logger, err)
					return nil, err
				}
				tokenValue := strings.Replace(token, TokenType, "", 1)
				access, err := store.GetByAccess(ctx, tokenValue)
				if err != nil {
					err = errors.New(6401, "Invalid Token", "未登录或登录已过期")
					logWithCtx(ctx, logger, err)
					return nil, err
				}
				if access == nil {
					err = errors.New(6401, UnAuthorized, "未登录或登录已过期")
					logWithCtx(ctx, logger, err)
					return nil, err
				}
				ctx = context.WithValue(ctx, UserId, access.GetUserID())
			default:
				log.Errorf("[authentication] message predicate err")
				return nil, errors.New(http.StatusBadRequest, "BadRequest", "参数格式错误")
			}
			return handler(ctx, req)
		}
	}
}
