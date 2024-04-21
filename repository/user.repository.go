package repository

import (
	"context"
	"reimbursement/repository/entity"
)

func (d *mysqlDatabase) CreateUser(ctx context.Context, data entity.User) error {
	query := d.Db.Model(data)
	query.Create(&data)

	return query.Error
}

func (d *mysqlDatabase) FindUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var data entity.User
	query := d.Db.Model(&data)
	query = query.Where("email = ?", email)
	query.First(&data)

	return data, query.Error
}
