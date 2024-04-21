package controllers

import (
	"net/http"
	"reimbursement/helper"
	hModels "reimbursement/helper/models"
	"reimbursement/usecase/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *controller) CreateEmployeeClaim(ctx *gin.Context) {
	helper.PrintHeader()

	var res hModels.Response
	var payload models.ReqCreateEmployeeClaim

	_, supportDocument, err := ctx.Request.FormFile("support_document")

	if err != nil {
		c.Logs.Println("Bad Request: ", err)
		res.Meta.Code = http.StatusBadRequest
		res.Meta.Message = "Bad Request"
		res.Meta.Title = "Failed"

		ctx.JSON(res.Meta.Code, res)
		return
	}

	if err := ctx.Bind(&payload); err != nil {
		c.Logs.Println("Bad Request: ", err)
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

	res = c.Usecase.CreateEmployeeClaim(ctx, payload, supportDocument)

	ctx.JSON(res.Meta.Code, res)
}

func (c *controller) GetAllEmployeeClaim(ctx *gin.Context) {
	helper.PrintHeader()

	var res hModels.Response
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	params := models.ParamsGetEmployeeClaim{
		Page:  page,
		Limit: limit,
	}
	res = c.Usecase.GetAllEmployeeClaim(ctx, params)

	ctx.JSON(res.Meta.Code, res)
}

func (c *controller) UpdateEmployeeClaim(ctx *gin.Context) {
	helper.PrintHeader()

	var res hModels.Response
	var payload models.ReqCreateEmployeeClaim
	id, _ := strconv.Atoi(ctx.Param("id"))

	_, supportDocument, err := ctx.Request.FormFile("support_document")

	if err != nil {
		c.Logs.Println("Bad Request: ", err)
		res.Meta.Code = http.StatusBadRequest
		res.Meta.Message = "Bad Request"
		res.Meta.Title = "Failed"

		ctx.JSON(res.Meta.Code, res)
		return
	}

	if err := ctx.Bind(&payload); err != nil {
		c.Logs.Println("Bad Request: ", err)
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

	res = c.Usecase.EditEmployeeClaim(ctx, id, payload, supportDocument)

	ctx.JSON(res.Meta.Code, res)
}

func (c *controller) DeleteEmployeeClaim(ctx *gin.Context) {
	var res hModels.Response

	id, _ := strconv.Atoi(ctx.Param("id"))

	res = c.Usecase.DeleteClaim(ctx, id)

	ctx.JSON(res.Meta.Code, res)
}

func (c *controller) GetAllEmployeeClaimAdmin(ctx *gin.Context) {
	helper.PrintHeader()

	var res hModels.Response
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	if limit == 0 {
		limit = 10
	}

	if page == 0 {
		page = 1
	}

	params := models.ParamsGetEmployeeClaim{
		Page:  page,
		Limit: limit,
	}
	res = c.Usecase.GetAllEmployeeClaimAdmin(ctx, params)

	ctx.JSON(res.Meta.Code, res)
}

func (c *controller) ApproveOrRejectClaim(ctx *gin.Context) {
	helper.PrintHeader()

	var res hModels.Response
	var payload models.ReqApprovaOrRejectClaim
	claimdId, _ := strconv.Atoi(ctx.Param("id"))

	if err := ctx.Bind(&payload); err != nil {
		c.Logs.Println("Bad Request: ", err)
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

	res = c.Usecase.ApproveOrRejectClaim(ctx, claimdId, payload)

	ctx.JSON(res.Meta.Code, res)
}
