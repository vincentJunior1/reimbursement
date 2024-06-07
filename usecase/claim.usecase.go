package usecase

import (
	"context"
	"mime/multipart"
	"reimbursement/constants"
	"reimbursement/helper"
	hModels "reimbursement/helper/models"
	"reimbursement/repository/entity"
	"reimbursement/usecase/models"
	"time"

	"github.com/gin-gonic/gin"
)

func (u *usecase) CreateEmployeeClaim(ctx *gin.Context, payload models.ReqCreateEmployeeClaim, supportDocument *multipart.FileHeader) hModels.Response {
	helper.PrintHeader()
	res := hModels.Response{}

	userInfo := &models.ClaimsJwtRes{}

	if ctx.Value("user") != nil {
		userInfo = ctx.Value("user").(*models.ClaimsJwtRes)
	}

	if userInfo.Id == 0 {
		u.Logs.Error("Token not valid")
		res.Meta = helper.MetaHelper(constants.Unauthorized)
		return res
	}

	fileName := ""
	if supportDocument != nil {
		fileName = supportDocument.Filename
		filePath := helper.GetEnv("FILEPATH") + helper.GetEnv("FILEPATH_CLAIM")

		if err := ctx.SaveUploadedFile(supportDocument, filePath); err != nil {
			res.Meta = helper.MetaHelper(constants.FailedSaveData)
			return res
		}

	}

	claimDate, _ := time.Parse("2006-01-02 15:04:05", payload.ClaimDate)
	data := entity.EmployeeClaim{
		EmployeeId:      userInfo.Id,
		CompanyId:       userInfo.CompanyId,
		ClaimCateogry:   payload.ClaimCateogry,
		ClaimDate:       claimDate,
		Currency:        payload.Currency,
		ClaimAmount:     payload.ClaimAmount,
		Status:          1,
		Description:     payload.Description,
		SupportDocument: fileName,
	}

	newData, err := u.DB.SaveEmployeeClaim(ctx, data)

	if err != nil {
		res.Meta = helper.MetaHelper(constants.FailedSaveData)
		return res
	}

	res.Meta = helper.MetaHelper(constants.SuccessCreateData)
	res.Data = newData
	return res
}

func (u *usecase) GetAllEmployeeClaim(ctx context.Context, params models.ParamsGetEmployeeClaim) hModels.Response {
	res := hModels.Response{}

	userInfo := &models.ClaimsJwtRes{}

	if ctx.Value("user") != nil {
		userInfo = ctx.Value("user").(*models.ClaimsJwtRes)
	}

	if userInfo.Id == 0 {
		u.Logs.Error("Token not valid", *userInfo)
		res.Meta = helper.MetaHelper(constants.Unauthorized)
		return res
	}

	data, pagination, err := u.DB.GetAllEmployeeClaim(ctx, userInfo.Id, params)

	if err != nil {
		u.Logs.Println("Error get data ", err)
		res.Meta = helper.MetaHelper(constants.ServerBusy)
		return res
	}

	res.Meta = helper.MetaHelper(constants.SuccesGetData)
	res.Data = data
	res.Page = &pagination
	return res
}

func (u *usecase) EditEmployeeClaim(ctx *gin.Context, id int, payload models.ReqCreateEmployeeClaim, supportDocument *multipart.FileHeader) hModels.Response {
	res := hModels.Response{}

	userInfo := &models.ClaimsJwtRes{}

	if ctx.Value("user") != nil {
		userInfo = ctx.Value("user").(*models.ClaimsJwtRes)
	}

	if userInfo.Id == 0 {
		u.Logs.Error("Token not valid")
		res.Meta = helper.MetaHelper(constants.Unauthorized)
		return res
	}

	_, err := u.DB.FindEmployeeClaim(ctx, id)

	if err != nil {
		u.Logs.Error("Data not found ", err)
		res.Meta = helper.MetaHelper(constants.DataNotFound)
		return res
	}

	fileName := supportDocument.Filename
	filePath := helper.GetEnv("FILEPATH") + helper.GetEnv("FILEPATH_CLAIM")

	if err := ctx.SaveUploadedFile(supportDocument, filePath); err != nil {
		res.Meta = helper.MetaHelper(constants.FailedSaveData)
		return res
	}

	claimDate, err := time.Parse("2006-01-02 15:04:05", payload.ClaimDate)
	data := &entity.EmployeeClaim{
		EmployeeId:      userInfo.Id,
		CompanyId:       userInfo.CompanyId,
		ClaimCateogry:   payload.ClaimCateogry,
		ClaimDate:       claimDate.Local(),
		Currency:        payload.Currency,
		ClaimAmount:     payload.ClaimAmount,
		Status:          1,
		Description:     payload.Description,
		SupportDocument: fileName,
		UpdatedAt:       time.Now(),
	}

	if err := u.DB.UpdatedEmployeeClaim(ctx, id, data); err != nil {
		res.Meta = helper.MetaHelper(constants.FailedSaveData)
		return res
	}

	res.Meta = helper.MetaHelper(constants.SuccessCreateData)
	res.Data = data
	return res
}

func (u *usecase) DeleteClaim(ctx context.Context, id int) hModels.Response {
	helper.PrintHeader()
	res := hModels.Response{}

	claim, err := u.DB.FindEmployeeClaim(ctx, id)

	if err != nil {
		u.Logs.Println("Data not found ")
		res.Meta = helper.MetaHelper(constants.DataNotFound)
		return res
	}

	if err := u.DB.DeleteEmployeeClaim(ctx, claim); err != nil {
		u.Logs.Println("Error delete employee claim: ", err)
		res.Meta = helper.MetaHelper(constants.FailedSaveData)
		return res
	}

	res.Meta = helper.MetaHelper(constants.SuccessAccepted)
	return res
}

func (u *usecase) GetAllEmployeeClaimAdmin(ctx context.Context, params models.ParamsGetEmployeeClaim) hModels.Response {
	res := hModels.Response{}

	userInfo := &models.ClaimsJwtRes{}

	if ctx.Value("user") != nil {
		userInfo = ctx.Value("user").(*models.ClaimsJwtRes)
	}

	if userInfo.Id == 0 {
		u.Logs.Error("Token not valid")
		res.Meta = helper.MetaHelper(constants.Unauthorized)
		return res
	}

	data, pagination, err := u.DB.GetAllEmployeeClaimAdmin(ctx, userInfo.CompanyId, params)

	if err != nil {
		u.Logs.Println("Error get data: ", err)
		res.Meta = helper.MetaHelper(constants.ServerBusy)
		return res
	}

	res.Meta = helper.MetaHelper(constants.SuccesGetData)
	res.Data = data
	res.Page = &pagination
	return res
}

func (u *usecase) ApproveOrRejectClaim(ctx context.Context, claimId int, payload models.ReqApprovaOrRejectClaim) hModels.Response {
	helper.PrintHeader()

	res := hModels.Response{}

	claim, err := u.DB.FindEmployeeClaim(ctx, claimId)

	if err != nil {
		u.Logs.Println("Data not found: ", err)
		res.Meta = helper.MetaHelper(constants.DataNotFound)
		return res
	}
	// status 1 pending
	// status 2 accepted
	// status 3 rejected
	claim.Status = payload.Status

	if err := u.DB.UpdatedEmployeeClaim(ctx, claimId, &claim); err != nil {
		u.Logs.Println("Failed Update Employee claim: ", err)
		res.Meta = helper.MetaHelper(constants.FailedSaveData)
		return res
	}

	res.Meta = helper.MetaHelper(constants.SuccessCreateData)
	return res
}
