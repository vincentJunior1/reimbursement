package usecase

import (
	"context"
	"net/http"
	hModels "reimbursement/helper/models"
)

func (u usecase) HealthCheck(ctx context.Context) hModels.Response {
	resp := hModels.Response{}

	resp.Meta.Code = http.StatusOK
	resp.Meta.Message = "Alive"
	return resp
}
