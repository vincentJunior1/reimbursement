package usecase

import (
	"context"
	"reimbursement/constants"
	"reimbursement/helper"
	hModels "reimbursement/helper/models"
	"reimbursement/repository/entity"
	"reimbursement/usecase/models"

	"github.com/golang-jwt/jwt/v5"

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

//

func (u *usecase) Login(ctx context.Context, payload models.ReqLogin) hModels.Response {
	res := hModels.Response{}

	user, err := u.DB.FindUserByEmail(ctx, payload.Email)

	if err != nil {
		u.Logs.Println("User not found: ", err)
		res.Meta = helper.MetaHelper(constants.DataNotFound)
		return res
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		res.Meta = helper.MetaHelper(constants.Unauthorized)
		res.Meta.Message = "Wrong Password!!!"
		return res
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         user.Id,
		"company_id": user.CompanyId,
		"name":       user.Name,
		"type":       user.Type,
	})

	tokenString, err := token.SignedString([]byte(helper.GetEnv("JWT_SECRET")))

	if err != nil {
		u.Logs.Println("Error Signed Jwt: ", err)
		res.Meta = helper.MetaHelper(constants.Unauthorized)
		return res
	}

	res.Meta = helper.MetaHelper(constants.SuccesGetData)
	res.Data = models.UserJwtRes{
		Token: tokenString,
	}
	return res
}
