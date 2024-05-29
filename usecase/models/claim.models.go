package models

type ReqCreateEmployeeClaim struct {
	ClaimCateogry *int    `form:"claim_category" json:"claim_category" validate:"required"`
	ClaimDate     string  `form:"claim_date" json:"claim_date" validate:"required"`
	Currency      string  `form:"currency" json:"currency" validate:"required"`
	ClaimAmount   float64 `form:"claim_amount" json:"claim_amount" validate:"required"`
	Description   string  `form:"description" json:"description"`
}

type ReqApprovaOrRejectClaim struct {
	Status int `json:"status" validate:"required"`
}
