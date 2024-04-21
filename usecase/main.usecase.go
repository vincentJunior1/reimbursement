package usecase

import (
	"context"
	"mime/multipart"
	hModels "reimbursement/helper/models"
	"reimbursement/repository"
	"reimbursement/usecase/models"

	"github.com/gin-gonic/gin"
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
		DeleteClaim(ctx context.Context, id int) hModels.Response
		Login(ctx context.Context, payload models.ReqLogin) hModels.Response
		CreateUser(ctx context.Context, payload models.ReqSaveUser) hModels.Response
		CreateCompany(ctx context.Context, payload models.ReqCompany) hModels.Response
		GetAllEmployeeClaim(ctx context.Context, params models.ParamsGetEmployeeClaim) hModels.Response
		GetAllEmployeeClaimAdmin(ctx context.Context, params models.ParamsGetEmployeeClaim) hModels.Response
		ApproveOrRejectClaim(ctx context.Context, claimId int, payload models.ReqApprovaOrRejectClaim) hModels.Response
		CreateEmployeeClaim(ctx *gin.Context, payload models.ReqCreateEmployeeClaim, supportDocument *multipart.FileHeader) hModels.Response
		EditEmployeeClaim(ctx *gin.Context, id int, payload models.ReqCreateEmployeeClaim, supportDocument *multipart.FileHeader) hModels.Response
	}
)

// InitializeV1Usecase ...
func InitializeV1Usecase(db repository.MysqlDatabase, l *logs.Logger) Usecase {
	return &usecase{
		Logs: l,
		DB:   db,
	}
}
