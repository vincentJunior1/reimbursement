package repository

import (
	"context"
	"reimbursement/repository/entity"
)

func (d *mysqlDatabase) CreateCompany(ctx context.Context, data *entity.Company) error {
	query := d.Db.Model(data)
	query = query.Create(data)

	query.Find(data)

	return query.Error
}

func (d *mysqlDatabase) FindCompanyById(ctx context.Context, id int) (entity.Company, error) {
	var data entity.Company
	query := d.Db.Model(data)
	query = query.Where("id = ?", id)
	query.First(&data)

	return data, query.Error
}
