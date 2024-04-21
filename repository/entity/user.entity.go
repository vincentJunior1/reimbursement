package entity

import "time"

type User struct {
	Id        int       `gorm:"column:id" json:"id"`
	Email     string    `gorm:"column:email" json:"email"`
	Password  string    `gorm:"column:password" json:"password"`
	Name      string    `gorm:"column:name" json:"name"`
	Address   string    `gorm:"column:address" json:"address"`
	Type      int       `gorm:"column:type" json:"type"`
	CompanyId int       `gorm:"column:company_id" json:"company_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) TypeUser() string {
	typeUser := map[int]string{
		1: "Admin",
		2: "Employee",
	}

	return typeUser[u.Type]
}
