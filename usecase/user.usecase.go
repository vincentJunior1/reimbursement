package usecase

import (
	"context"
	"reimbursement/constants"
	"reimbursement/helper"
	hModels "reimbursement/helper/models"
	"reimbursement/repository/entity"
	"reimbursement/usecase/models"

	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) CreateUser(ctx context.Context, payload models.ReqSaveUser) hModels.Response {
	res := hModels.Response{}

	company, err := u.DB.FindCompanyById(ctx, payload.CompanyId)

	if err != nil {
		u.Logs.Println("Company Not Found")
		res.Meta = helper.MetaHelper(constants.DataNotFound)
		return res
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), 14)
	data := entity.User{
		Name:      payload.Name,
		Email:     payload.Email,
		CompanyId: company.Id,
		Password:  string(password),
		Address:   payload.Email,
		Type:      payload.Type,
	}

	if err := u.DB.CreateUser(ctx, data); err != nil {
		u.Logs.Println("Failed Create User: ", err)
		res.Meta = helper.MetaHelper(constants.FailedSaveData)
		return res
	}

	res.Meta = helper.MetaHelper(constants.SuccessCreateData)
	return res

}
