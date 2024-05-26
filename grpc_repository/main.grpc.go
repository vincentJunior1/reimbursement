package grpcrepository

import (
	"context"
	pb "reimbursement/proto-go"
)

type (
	server struct {
	}

	Server interface {
		SayHiUser(ctx context.Context, user *pb.User) *pb.User
	}
)
