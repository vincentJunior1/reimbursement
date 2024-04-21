package models

// CheckServiceReq ..
type CheckServiceReq struct {
	ServiceName string
	Host        string
	Endpoint    string
	Method      string
}

// ResponseMessage ..
type ResponseMessage struct {
	ResponseCode int               `json:"rc"`
	Message      string            `json:"message"`
	Data         []DataHealthCheck `json:"data"`
}

// DataHealthCheck ..
type DataHealthCheck struct {
	ServiceName    string `json:"serviceName"`
	Host           string `json:"host"`
	StatusCode     int    `json:"statusCode"`
	AdditionalData string `json:"additionalData"`
}
