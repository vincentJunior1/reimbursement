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
		CreateUser(ctx *gin.Context)
		HealthCheck(ctx *gin.Context)
		CreateCompany(ctx *gin.Context)
	}
)

// InitializeController ..
func InitializeV1Controller(uc usecase.Usecase, l *logs.Logger) Controller {
	return &controller{
		Logs:    l,
		Usecase: uc,
	}
}
