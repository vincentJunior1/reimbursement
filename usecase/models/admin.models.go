package models

type ReqSaveUser struct {
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	CompanyId int    `json:"company_id" validate:"required"`
	Type      int    `json:"type" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Address   string `json:"address" validate:"required"`
}
