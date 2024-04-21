package middleware

import (
	"net/http"
	"reimbursement/constants"
	"reimbursement/helper"
	hModels "reimbursement/helper/models"
	"reimbursement/usecase/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type (
	middleware struct {
		Log *logrus.Logger
	}

	Middleware interface {
		JwtAuthorization() gin.HandlerFunc
	}
)

func InitMiddleware(log *logrus.Logger) Middleware {
	return &middleware{
		Log: log,
	}
}

func (m *middleware) JwtAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		newData := &models.ClaimsJwtRes{}
		token, err := jwt.ParseWithClaims(tokenString, newData, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unexpected signing method"})
				c.Abort()
			}
			return []byte(helper.GetEnv("JWT_SECRET")), nil
		})

		if err != nil {
			logrus.Error("Error token :", err)
			res := hModels.Response{
				Meta: helper.MetaHelper(constants.Unauthorized),
			}
			c.JSON(res.Meta.Code, res)
		}
		logrus.Println(token)
		c.Set("user", newData)
		c.Next()
	}
}
