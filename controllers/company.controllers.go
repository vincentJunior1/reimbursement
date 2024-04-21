package controllers

import (
	"context"
	"fmt"
	"net/http"
	"reimbursement/constants"
	"reimbursement/helper"
	hModels "reimbursement/helper/models"
	"reimbursement/usecase/models"
	"time"

	"github.com/gin-gonic/gin"
)

func (c *controller) CreateCompany(ctx *gin.Context) {
	helper.PrintHeader()

	var res hModels.Response
	var payload models.ReqCompany

	if err := ctx.BindJSON(&payload); err != nil {
		res.Meta.Code = http.StatusBadRequest
		res.Meta.Message = "Bad Request"
		res.Meta.Title = "Failed"

		ctx.JSON(res.Meta.Code, res)
		return
	}

	if err := helper.Validatestruct(payload); err != nil {
		res.Meta.Code = http.StatusBadRequest
		res.Meta.Message = err.Error()
		res.Meta.Title = "Failed"

		ctx.JSON(res.Meta.Code, res)
		return
	}

	ctxz := context.WithValue(context.Background(), constants.SPAN_ID, fmt.Sprintf("%d-%d", constants.SPAN_ID_LOGIN, time.Now().UnixNano()))
	res = c.Usecase.CreateCompany(ctxz, payload)

	ctx.JSON(res.Meta.Code, res)
}
