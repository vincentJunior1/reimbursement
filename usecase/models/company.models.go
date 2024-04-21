package models

type ReqCompany struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
}
