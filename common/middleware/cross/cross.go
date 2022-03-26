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
			tr.ReplyHeader().Set("Access-Control-Allow-Credentials", "true")
			tr.ReplyHeader().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			tr.ReplyHeader().Set("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			tr.ReplyHeader().Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
			return handler(ctx, req)
		}
	}
}
