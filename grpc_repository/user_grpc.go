package grpcrepository

import (
	"context"
	pb "reimbursement/proto-go"
)

func (s *server) SayHiUser(ctx context.Context, user *pb.User) *pb.User {
	return &pb.User{
		Id:   user.Id,
		Name: "Biji Goreng" + user.Name,
	}
}
