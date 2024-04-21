package helper

import (
	"fmt"
	"net/http"
	"os"
	hModels "reimbursement/helper/models"
	"runtime"
)

// GetEnv ..
func GetEnv(key string) string {

	// load .env file

	// if err != nil {
	// 	os.Setenv(key, defaultString)
	// }

	value := os.Getenv(key)
	// if value == "" {
	// 	logrus.Fatalf("ENV %v not found", key)
	// }

	return value
	// viper.SetConfigType("env")
	// viper.AddConfigPath(".")
	// viper.SetConfigName(".env")
	// viper.AutomaticEnv()
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	logrus.Fatalf("Error while reading config file %s", err)
	// }
	// value, ok := viper.Get(key).(string)
	// if !ok {
	// 	logrus.Fatalf("Invalid type assertion %s", key)
	// }

	// return value
}

func GetCaller() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func PrintHeader() {
	pc, _, _, _ := runtime.Caller(1)
	fmt.Printf("<======> %s <======>", runtime.FuncForPC(pc).Name())
	fmt.Println()
}

func MetaHelper(code int) hModels.MetaData {
	resp := hModels.MetaData{}
	switch code {
	case http.StatusCreated:
		resp.Code = http.StatusCreated
		resp.Message = "Success Save Data"
		resp.Title = "Success"
	case http.StatusOK:
		resp.Code = http.StatusOK
		resp.Message = "Success Get Data"
		resp.Title = "Success"
	case http.StatusNotFound:
		resp.Code = http.StatusNotFound
		resp.Message = "Data not found"
		resp.Title = "Failed"
	case http.StatusUnprocessableEntity:
		resp.Code = http.StatusUnprocessableEntity
		resp.Message = "Failed Create Data"
		resp.Title = "Failed"
	case http.StatusUnauthorized:
		resp.Code = http.StatusUnauthorized
		resp.Message = "Failed Login"
		resp.Title = "Unauthorized"
	case 500:
		resp.Code = 500
		resp.Message = "Something Wrong"
		resp.Title = "Failed"
	case http.StatusAccepted:
		resp.Code = http.StatusAccepted
		resp.Message = "Success"
		resp.Title = "Success"
	case http.StatusForbidden:
		resp.Code = http.StatusForbidden
		resp.Message = "You're not allowed"
		resp.Title = "Failed"
	}

	return resp
}
