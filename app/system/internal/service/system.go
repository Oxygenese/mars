package service

import (
	"context"
	"fmt"
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
	fmt.Println("[system service] request: ", req)
	var err error
	msg := &api.Message{
		Request: req,
		Context: ctx,
	}
	//if !s.taskManager.IsExecutorExists(req.Operate) {
	//	err = errors.New(404, "Not Found", "NotFound")
	//	log.Errorf("[service] PushMessage err :%s", err)
	//	return nil, err
	//}
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
