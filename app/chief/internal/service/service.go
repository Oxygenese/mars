package service

import (
	"context"
	"github.com/google/wire"
	pb "github.com/mars-projects/mars/api/chief"
	"github.com/mars-projects/mars/app/chief/internal/task"
	"github.com/mars-projects/mars/common/transaction"
)

// ProviderServiceSet is service providers.
var ProviderServiceSet = wire.NewSet(NewChiefService)

type ChiefService struct {
	pb.UnimplementedChiefServer
	taskManager *task.TasksManager
}

func NewChiefService(manager *task.TasksManager) *ChiefService {
	return &ChiefService{taskManager: manager}
}

type Image struct {
	Path string
	Cap  string
}

func (s *ChiefService) OnMessageReceived(ctx context.Context, req *pb.Message) (*pb.Reply, error) {
	msg := transaction.NewTransJsonMessage(req)
	err := s.taskManager.PushMessage(msg)
	if err != nil {
		return nil, err
	}
	return &pb.Reply{Code: 200, Message: "发送成功", Success: true}, nil
}
