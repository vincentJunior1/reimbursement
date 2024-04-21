package usecase

import (
	"context"
	"reimbursement/constants"
	"reimbursement/helper"
	hModels "reimbursement/helper/models"
	"reimbursement/repository/entity"
	"reimbursement/usecase/models"
)

func (u *usecase) CreateCompany(ctx context.Context, payload models.ReqCompany) hModels.Response {
	helper.PrintHeader()
	res := hModels.Response{}

	data := &entity.Company{
		Name:    payload.Name,
		Address: payload.Address,
	}

	if err := u.DB.CreateCompany(ctx, data); err != nil {
		u.Logs.Error("Error create company: ", err)
		res.Meta = helper.MetaHelper(constants.FailedSaveData)
		return res
	}

	res.Meta = helper.MetaHelper(constants.SuccessCreateData)

	res.Data = data
	return res

}
