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
