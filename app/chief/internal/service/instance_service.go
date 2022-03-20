package service

import (
	"context"

	pb "github.com/mars-projects/mars/api/chief"
)

type InstanceServiceService struct {
	pb.UnimplementedInstanceServiceServer
}

func NewInstanceServiceService() *InstanceServiceService {
	return &InstanceServiceService{}
}

func (s *InstanceServiceService) OnMessageReceived(ctx context.Context, req *pb.Message) (*pb.Reply, error) {
	return &pb.Reply{}, nil
}
