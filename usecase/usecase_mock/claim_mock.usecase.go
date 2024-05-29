// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecase/main.usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	multipart "mime/multipart"
	reflect "reflect"
	models "reimbursement/helper/models"
	models0 "reimbursement/usecase/models"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockUsecase is a mock of Usecase interface.
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase.
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance.
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// ApproveOrRejectClaim mocks base method.
func (m *MockUsecase) ApproveOrRejectClaim(ctx context.Context, claimId int, payload models0.ReqApprovaOrRejectClaim) models.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApproveOrRejectClaim", ctx, claimId, payload)
	ret0, _ := ret[0].(models.Response)
	return ret0
}

// ApproveOrRejectClaim indicates an expected call of ApproveOrRejectClaim.
func (mr *MockUsecaseMockRecorder) ApproveOrRejectClaim(ctx, claimId, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApproveOrRejectClaim", reflect.TypeOf((*MockUsecase)(nil).ApproveOrRejectClaim), ctx, claimId, payload)
}

// CreateCompany mocks base method.
func (m *MockUsecase) CreateCompany(ctx context.Context, payload models0.ReqCompany) models.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCompany", ctx, payload)
	ret0, _ := ret[0].(models.Response)
	return ret0
}

// CreateCompany indicates an expected call of CreateCompany.
func (mr *MockUsecaseMockRecorder) CreateCompany(ctx, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompany", reflect.TypeOf((*MockUsecase)(nil).CreateCompany), ctx, payload)
}

// CreateEmployeeClaim mocks base method.
func (m *MockUsecase) CreateEmployeeClaim(ctx *gin.Context, payload models0.ReqCreateEmployeeClaim, supportDocument *multipart.FileHeader) models.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmployeeClaim", ctx, payload, supportDocument)
	ret0, _ := ret[0].(models.Response)
	return ret0
}

// CreateEmployeeClaim indicates an expected call of CreateEmployeeClaim.
func (mr *MockUsecaseMockRecorder) CreateEmployeeClaim(ctx, payload, supportDocument interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmployeeClaim", reflect.TypeOf((*MockUsecase)(nil).CreateEmployeeClaim), ctx, payload, supportDocument)
}

// CreateUser mocks base method.
func (m *MockUsecase) CreateUser(ctx context.Context, payload models0.ReqSaveUser) models.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, payload)
	ret0, _ := ret[0].(models.Response)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUsecaseMockRecorder) CreateUser(ctx, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUsecase)(nil).CreateUser), ctx, payload)
}

// DeleteClaim mocks base method.
func (m *MockUsecase) DeleteClaim(ctx context.Context, id int) models.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteClaim", ctx, id)
	ret0, _ := ret[0].(models.Response)
	return ret0
}

// DeleteClaim indicates an expected call of DeleteClaim.
func (mr *MockUsecaseMockRecorder) DeleteClaim(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteClaim", reflect.TypeOf((*MockUsecase)(nil).DeleteClaim), ctx, id)
}

// EditEmployeeClaim mocks base method.
func (m *MockUsecase) EditEmployeeClaim(ctx *gin.Context, id int, payload models0.ReqCreateEmployeeClaim, supportDocument *multipart.FileHeader) models.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditEmployeeClaim", ctx, id, payload, supportDocument)
	ret0, _ := ret[0].(models.Response)
	return ret0
}

// EditEmployeeClaim indicates an expected call of EditEmployeeClaim.
func (mr *MockUsecaseMockRecorder) EditEmployeeClaim(ctx, id, payload, supportDocument interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditEmployeeClaim", reflect.TypeOf((*MockUsecase)(nil).EditEmployeeClaim), ctx, id, payload, supportDocument)
}

// GetAllEmployeeClaim mocks base method.
func (m *MockUsecase) GetAllEmployeeClaim(ctx context.Context, params models0.ParamsGetEmployeeClaim) models.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEmployeeClaim", ctx, params)
	ret0, _ := ret[0].(models.Response)
	return ret0
}

// GetAllEmployeeClaim indicates an expected call of GetAllEmployeeClaim.
func (mr *MockUsecaseMockRecorder) GetAllEmployeeClaim(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEmployeeClaim", reflect.TypeOf((*MockUsecase)(nil).GetAllEmployeeClaim), ctx, params)
}

// GetAllEmployeeClaimAdmin mocks base method.
func (m *MockUsecase) GetAllEmployeeClaimAdmin(ctx context.Context, params models0.ParamsGetEmployeeClaim) models.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEmployeeClaimAdmin", ctx, params)
	ret0, _ := ret[0].(models.Response)
	return ret0
}

// GetAllEmployeeClaimAdmin indicates an expected call of GetAllEmployeeClaimAdmin.
func (mr *MockUsecaseMockRecorder) GetAllEmployeeClaimAdmin(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEmployeeClaimAdmin", reflect.TypeOf((*MockUsecase)(nil).GetAllEmployeeClaimAdmin), ctx, params)
}

// HealthCheck mocks base method.
func (m *MockUsecase) HealthCheck(ctx context.Context) models.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck", ctx)
	ret0, _ := ret[0].(models.Response)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockUsecaseMockRecorder) HealthCheck(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockUsecase)(nil).HealthCheck), ctx)
}

// Login mocks base method.
func (m *MockUsecase) Login(ctx context.Context, payload models0.ReqLogin) models.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, payload)
	ret0, _ := ret[0].(models.Response)
	return ret0
}

// Login indicates an expected call of Login.
func (mr *MockUsecaseMockRecorder) Login(ctx, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUsecase)(nil).Login), ctx, payload)
}
