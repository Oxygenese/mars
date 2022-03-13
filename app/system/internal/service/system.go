package service

import (
	"context"

	pb "github.com/mars-projects/mars/api/system"
)

type SystemService struct {
	pb.UnimplementedSystemServer
}

func NewSystemService() *SystemService {
	return &SystemService{}
}

func (s *SystemService) SysUserInfo(ctx context.Context, req *pb.SysUserInfoReq) (*pb.SysUserReply, error) {
	return &pb.SysUserReply{}, nil
}
