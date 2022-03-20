package service

import (
	"context"
	pb "github.com/mars-projects/mars/api/chief"
	"github.com/mars-projects/mars/app/chief/internal/task"
	"github.com/mars-projects/mars/common/framework"
	"log"
)

type ImageService struct {
	pb.UnimplementedImageServiceServer
	taskManager *task.TaskManager
}

func NewImageService(manager *task.TaskManager) *ImageService {
	if err := manager.Start(); err != nil {
		panic(err)
	}
	return &ImageService{
		taskManager: manager,
	}
}

func (s *ImageService) OnMessageReceived(ctx context.Context, req *pb.Message) (*pb.Reply, error) {
	msg := framework.NewTransJsonMessage(req)
	if targetSession := msg.GetToSession(); targetSession != 0 {
		if err := s.taskManager.PushMessage(msg); err != nil {
			log.Printf("<image> push message [%08X] from %s to session [%08X] fail: %s", msg.GetID(), msg.GetSender(), targetSession, err.Error())
		}
		return nil, nil
	}
	var err error
	switch msg.GetID() {
	case framework.QueryDiskImageRequest:
	case framework.GetDiskImageRequest:
	case framework.CreateDiskImageRequest:
	case framework.DeleteDiskImageRequest:
	case framework.ModifyDiskImageRequest:
	case framework.SynchronizeDiskImageRequest:

	case framework.QueryMediaImageRequest:
	case framework.GetMediaImageRequest:
	case framework.CreateMediaImageRequest:
	case framework.DeleteMediaImageRequest:
	case framework.ModifyMediaImageRequest:
	case framework.SynchronizeMediaImageRequest:
	case framework.DiskImageUpdatedEvent:
	default:
		log.Printf("<image> message [%08X] from %s.[%08X] ignored", msg.GetID(), msg.GetSender(), msg.GetFromSession())
		return &pb.Reply{Code: 200, Message: "请求已忽略"}, nil
	}

	//Invoke transaction
	err = s.taskManager.InvokeTask(msg)
	if err != nil {
		log.Printf("<image> invoke transaction with message [%08X] fail: %s", msg.GetID(), err.Error())
	}
	return &pb.Reply{Code: 200, Message: "请求成功"}, nil
}
