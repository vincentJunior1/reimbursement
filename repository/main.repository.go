package repository

import (
	"context"
	"fmt"
	"reimbursement/helper"
	"reimbursement/repository/entity"

	// uModels "reimbursement/usecase/models"

	logs "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// MysqlDatabase ...
type (
	mysqlDatabase struct {
		Db   *gorm.DB
		Logs *logs.Logger
	}

	// interface from MysqlDatabase
	MysqlDatabase interface {
		FindCompanyById(ctx context.Context, id int) (entity.Company, error)
		CreateUser(ctx context.Context, data entity.User) error
		CreateCompany(ctx context.Context, data *entity.Company) error
	}
)

// InitializeMysqlDatabase ..
func InitializeMysqlDatabase(log *logs.Logger) MysqlDatabase {
	return &mysqlDatabase{
		Db:   ConnectMysql(),
		Logs: log,
	}
}

// logMode ...
var logMode = map[string]logger.LogLevel{
	"silent": logger.Silent,
	"error":  logger.Error,
	"warn":   logger.Warn,
	"info":   logger.Info,
}

// ConnectMysql connection ...
func ConnectMysql() *gorm.DB {
	username := helper.GetEnv("USER_MYSQL")
	password := helper.GetEnv("PASS_MYSQL")
	host := helper.GetEnv("HOST_MYSQL")
	port := helper.GetEnv("PORT_MYSQL")
	dbName := helper.GetEnv("DB_MYSQL")
	debug := helper.GetEnv("DEBUG_MYSQL")
	mode := helper.GetEnv("LOG_MODE_MYSQL")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logMode[mode]),
	})
	if err != nil {
		logs.WithFields(logs.Fields{"Message": err}).Error(helper.GetCaller())
		panic("Error open mysql connection")
	}

	logs.Info("Mysql connected successfully")
	if debug == "true" {
		return db.Debug()
	}

	return db
}
