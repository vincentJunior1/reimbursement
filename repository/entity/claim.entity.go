package entity

import "time"

type EmployeeClaim struct {
	Id              int       `gorm:"column:id" json:"id"`
	EmployeeId      int       `gorm:"column:employee_id" json:"employee_id"`
	CompanyId       int       `gorm:"column:company_id" json:"company_id"`
	ClaimCateogry   int       `gorm:"column:claim_category" json:"claim_category"`
	ClaimDate       time.Time `gorm:"column:claim_date" json:"claim_date"`
	Currency        string    `gorm:"column:currency" json:"currency"`
	ClaimAmount     float64   `gorm:"column:claim_amount" json:"claim_amount"`
	Status          int       `gorm:"column:status" json:"status"`
	Description     string    `gorm:"column:description" json:"description"`
	SupportDocument string    `gorm:"column:support_document" json:"support_document"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (EmployeeClaim) TableName() string {
	return "employee_claims"
}
