package usecase

import (
	"context"
	hModels "reimbursement/helper/models"
	"reimbursement/repository"
	"reimbursement/usecase/models"

	logs "github.com/sirupsen/logrus"
)

type (
	usecase struct {
		DB repository.MysqlDatabase
		// Host hosts.Host
		Logs *logs.Logger
	}

	Usecase interface {
		HealthCheck(ctx context.Context) hModels.Response
		CreateUser(ctx context.Context, payload models.ReqSaveUser) hModels.Response
		CreateCompany(ctx context.Context, payload models.ReqCompany) hModels.Response
	}
)

// InitializeV1Usecase ...
func InitializeV1Usecase(db repository.MysqlDatabase, l *logs.Logger) Usecase {
	return &usecase{
		Logs: l,
		DB:   db,
	}
}
