package repository

import (
	"context"
	hModels "reimbursement/helper/models"
	"reimbursement/repository/entity"
	"reimbursement/usecase/models"
)

func (d *mysqlDatabase) SaveEmployeeClaim(ctx context.Context, data entity.EmployeeClaim) (entity.EmployeeClaim, error) {
	query := d.Db.Model(&data)
	query = query.Save(data)

	query.First(&data)

	return data, query.Error
}

func (d *mysqlDatabase) GetAllEmployeeClaim(ctx context.Context, employeeId int, params models.ParamsGetEmployeeClaim) ([]entity.EmployeeClaim, hModels.Page, error) {
	var data []entity.EmployeeClaim
	var pagination hModels.Page
	query := d.Db.Model(&data)
	query = query.Where("employee_id = ?", employeeId)
	query.Count(&pagination.TotalData)
	offset := int(pagination.TotalData / int64(params.Limit))
	pagination.Page = params.Page
	query = query.Limit(params.Limit).Offset(offset - 1)

	query.Find(&data)

	return data, pagination, query.Error
}

func (d *mysqlDatabase) FindEmployeeClaim(ctx context.Context, id int) (entity.EmployeeClaim, error) {
	var data entity.EmployeeClaim

	query := d.Db.Model(&data)
	query = query.Where("id = ?", id)
	query.First(&data)

	return data, query.Error
}

func (d *mysqlDatabase) UpdatedEmployeeClaim(ctx context.Context, id int, data *entity.EmployeeClaim) error {
	query := d.Db.Model(data)
	query = query.Where("id = ?", id)
	query = query.Updates(data)

	query.First(data)

	return query.Error
}

func (d *mysqlDatabase) DeleteEmployeeClaim(ctx context.Context, data entity.EmployeeClaim) error {
	query := d.Db.Model(data)
	query.Delete(&data)

	return query.Error
}

func (d *mysqlDatabase) GetAllEmployeeClaimAdmin(ctx context.Context, companyId int, params models.ParamsGetEmployeeClaim) ([]entity.EmployeeClaim, hModels.Page, error) {
	var data []entity.EmployeeClaim
	var pagination hModels.Page
	query := d.Db.Model(&data)
	query = query.Where("company_id = ?", companyId)
	query.Count(&pagination.TotalData)
	offset := int(pagination.TotalData / int64(params.Limit))
	pagination.Page = params.Page
	query = query.Limit(params.Limit).Offset(offset - 1)

	query.Find(&data)

	return data, pagination, query.Error
}
