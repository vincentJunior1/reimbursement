package controllers

import (
	"github.com/gin-gonic/gin"
)

func (c *controller) HealthCheck(ctx *gin.Context) {
	resp := c.Usecase.HealthCheck(ctx)

	ctx.JSON(200, resp)
}
