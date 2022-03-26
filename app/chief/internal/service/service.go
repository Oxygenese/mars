package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/mars-projects/mars/api"
	"github.com/mars-projects/mars/app/chief/internal/task"
)

// ProviderServiceSet is service providers.
var ProviderServiceSet = wire.NewSet(NewChiefService)

type ChiefService struct {
	api.UnimplementedChiefServer
	taskManager *task.TasksManager
}

func NewChiefService(manager *task.TasksManager) *ChiefService {
	return &ChiefService{taskManager: manager}
}

type Image struct {
	Path string
	Cap  string
}

func (s *ChiefService) OnMessageReceived(ctx context.Context, req *api.Request) (*api.Reply, error) {
	var err error
	msg := &api.Message{
		Request: req,
		Context: ctx,
	}
	if !s.taskManager.IsExecutorExists(req.GetOperate()) {
		err = errors.New(404, "Not Found", "NotFound")
		log.Errorf("[service] PushMessage err :%s", err)
		return nil, err
	}
	err = s.taskManager.PushMessage(msg)
	if err != nil {
		log.Errorf("[service] PushMessage err :%s", err)
		return nil, err
	}
	if !s.taskManager.IsSync(req.GetOperate()) {
		res := <-s.taskManager.GetResChan(req.GetRequestId())
		return res, nil
	}
	return &api.Reply{
		Code:    200,
		Message: "发送成功",
	}, nil
}
