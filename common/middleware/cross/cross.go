package cross

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"net/http"
)

func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, errors.New(http.StatusBadRequest, "BadRequest", "")
			}
			tr.ReplyHeader().Set("Access-Control-Allow-Origin", "*")
			tr.ReplyHeader().Set("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
			tr.ReplyHeader().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
			tr.ReplyHeader().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			tr.ReplyHeader().Set("Access-Control-Allow-Credentials", "true")
			return handler(ctx, req)
		}
	}
}
