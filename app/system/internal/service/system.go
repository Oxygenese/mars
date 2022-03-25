package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/mars-projects/mars/api"
	"github.com/mars-projects/mars/app/system/internal/task"
)

// ProviderServiceSet is service providers.
var ProviderServiceSet = wire.NewSet(NewSystemService)

func NewSystemService(manager *task.TasksManager) *SystemService {
	return &SystemService{taskManager: manager}
}

type SystemService struct {
	api.UnimplementedSystemServer
	taskManager *task.TasksManager
}

func (s *SystemService) OnMessageReceived(ctx context.Context, req *api.Request) (*api.Reply, error) {
	msg := &api.Message{
		Request: req,
		Context: ctx,
	}
	err := s.taskManager.PushMessage(msg)
	if err != nil {
		log.Errorf("[service] PushMessage err :%s", err)
		return nil, err
	}
	if !s.taskManager.IsSync(req.GetOperation()) {
		res := <-s.taskManager.GetResChan(req.GetRequestId())
		return res, nil
	}
	return &api.Reply{
		Code:    200,
		Message: "发送成功",
	}, nil
}
