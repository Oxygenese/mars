package service

import (
	"context"
	"github.com/mars-projects/mars/app/system/internal/biz"

	pb "github.com/mars-projects/mars/api/system"
)

type SystemService struct {
	pb.UnimplementedSystemServer
	biz *biz.SysUser
}

func NewSystemService(option *biz.BizsOption) *SystemService {
	return &SystemService{biz: option.SysUser}
}

func (s *SystemService) SysUserInfo(ctx context.Context, req *pb.SysUserInfoReq) (*pb.SysUserReply, error) {
	var reply pb.SysUserReply
	s.biz.FindByUsername(req, &reply)
	return &reply, nil
}
