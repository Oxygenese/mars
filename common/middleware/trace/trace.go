package trace

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/google/uuid"
	"github.com/mars-projects/mars/api"
	"net/http"
)

func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			switch req.(type) {
			case *api.Request:
				message := req.(*api.Request)
				if message.RequestId == "" {
					newUUID, err := uuid.NewUUID()
					if err != nil {
						return nil, err
					}
					message.RequestId = newUUID.String()
				}
			default:
				return nil, errors.New(http.StatusBadRequest, "BadRequest", "参数错误")
			}
			return handler(ctx, req)
		}
	}
}
