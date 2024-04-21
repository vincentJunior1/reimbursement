package controllers

import (
	"reimbursement/usecase"

	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

//go:generate mockery --name Controller --output ../mocks

// type key string

// Controller ...
type (
	controller struct {
		Usecase usecase.Usecase
		Logs    *logs.Logger
	}
	Controller interface {
		Login(ctx *gin.Context)
		CreateUser(ctx *gin.Context)
		HealthCheck(ctx *gin.Context)
		CreateCompany(ctx *gin.Context)
		GetAllEmployeeClaim(ctx *gin.Context)
		CreateEmployeeClaim(ctx *gin.Context)
		UpdateEmployeeClaim(ctx *gin.Context)
		DeleteEmployeeClaim(ctx *gin.Context)
		ApproveOrRejectClaim(ctx *gin.Context)
		GetAllEmployeeClaimAdmin(ctx *gin.Context)
	}
)

// InitializeController ..
func InitializeV1Controller(uc usecase.Usecase, l *logs.Logger) Controller {
	return &controller{
		Logs:    l,
		Usecase: uc,
	}
}
