package usecase_test

import (
	"errors"
	"net/http/httptest"
	"reimbursement/helper"
	hModels "reimbursement/helper/models"
	entity "reimbursement/repository/entity"
	repoMock "reimbursement/repository/repository_mock"
	usecase "reimbursement/usecase"
	"reimbursement/usecase/models"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCreateEmployeeData(t *testing.T) {
	ctrl := gomock.NewController(t)
	gin.SetMode(gin.TestMode)
	ctx := gin.CreateTestContextOnly(httptest.NewRecorder(), &gin.Engine{})

	userClaims := models.ClaimsJwtRes{
		Id:        1,
		CompanyId: 1,
		Name:      "Vincent",
		Type:      1,
	}
	ctx.Set("user", &userClaims)
	defer ctrl.Finish()
	claimCategory := helper.ToIntPointer(1)
	repoMock := repoMock.NewMockMysqlDatabase(ctrl)
	mockEmployee := usecase.InitializeV1Usecase(repoMock, logrus.New())
	employeeModels := []struct {
		err   error
		args  entity.EmployeeClaim
		input models.ReqCreateEmployeeClaim
		got   hModels.Response
	}{
		{
			err: nil,
			args: entity.EmployeeClaim{
				Id:              0,
				EmployeeId:      1,
				CompanyId:       1,
				ClaimCateogry:   claimCategory,
				Currency:        "",
				ClaimAmount:     64000.00,
				ClaimDate:       time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
				Status:          1,
				Description:     "",
				SupportDocument: "",
				CreatedAt:       time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:       time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			input: models.ReqCreateEmployeeClaim{
				ClaimCateogry: claimCategory,
				ClaimDate:     time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).String(),
				ClaimAmount:   64000.00,
			},
			got: hModels.Response{
				Data: nil,
				Meta: helper.MetaHelper(201),
			},
		},
		{
			err: errors.New("database error"),
			args: entity.EmployeeClaim{
				Id:              0,
				EmployeeId:      1,
				CompanyId:       1,
				ClaimCateogry:   nil,
				ClaimDate:       time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
				Currency:        "",
				ClaimAmount:     64000.00,
				Status:          1,
				Description:     "",
				SupportDocument: "",
				CreatedAt:       time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt:       time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			input: models.ReqCreateEmployeeClaim{
				ClaimCateogry: nil,
				ClaimDate:     time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).String(),
				ClaimAmount:   64000.00,
			},
			got: hModels.Response{
				Data: nil,
				Meta: helper.MetaHelper(422),
			},
		},
	}
	for _, val := range employeeModels {
		repoMock.EXPECT().SaveEmployeeClaim(ctx, val.args).Return(val.args, val.err).AnyTimes()
		saveEmployeeClaim := mockEmployee.CreateEmployeeClaim(ctx, val.input, nil)
		assert.Equal(t, val.got.Meta, saveEmployeeClaim.Meta)
	}
}
