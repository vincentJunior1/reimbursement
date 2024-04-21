package models

import "github.com/golang-jwt/jwt/v5"

type ReqSaveUser struct {
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	CompanyId int    `json:"company_id" validate:"required"`
	Type      int    `json:"type" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Address   string `json:"address" validate:"required"`
}

type ReqLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserJwtRes struct {
	Token string `json:"token"`
}

type ClaimsJwtRes struct {
	jwt.RegisteredClaims
	Id        int    `json:"id"`
	CompanyId int    `json:"company_id"`
	Name      string `json:"name"`
	Type      int    `json:"type"`
}
