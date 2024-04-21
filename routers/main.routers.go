package routers

import (
	"time"

	"reimbursement/controllers"
	"reimbursement/helper"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

type (
	RouterInterface interface {
		StartServer() error
		routerControllers()
	}

	Router struct {
		address    string
		port       string
		Gin        *gin.Engine
		Logs       *logs.Logger
		Path       map[string]string
		Controller controllers.Controller
	}
)

// InitializeRouter return sturct that will implement the abs. class
func InitializeRouter(ctrl controllers.Controller, l *logs.Logger) RouterInterface {
	gin.SetMode(helper.GetEnv("ROUTER_SETMODE"))
	return &Router{
		address: helper.GetEnv("ROUTER_SERVER_ADDRESS"),
		port:    helper.GetEnv("ROUTER_PORT"),
		Path: map[string]string{
			"PATH_VERSION": helper.GetEnv("PATH_VERSION"),
			"PATH_MAIN":    helper.GetEnv("PATH_MAIN"),
		},
		Gin:        gin.New(),
		Logs:       l,
		Controller: ctrl,
	}
}

func (r *Router) StartServer() error {
	r.Logs.Info("Starting Server on ", r.address+r.port)
	r.routerControllers()
	// err := r.Gin.Run(r.port)

	err := helper.GinServerUp(r.address+r.port, r.Gin)
	if err != nil {
		r.Logs.Error("[GinServerUp]Error: ", err)
		return err
	}

	return nil
}

// routerControllers ...
func (r *Router) routerControllers() {

	r.Gin.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "PATCH", "PUT", "POST", "HEAD", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization", "access-control-allow-origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		//AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge: 12 * time.Hour,
	}))

	// }

	main := r.Gin.Group(r.Path["PATH_MAIN"])
	{
		// version 1
		v1 := main.Group("/v1")
		company := v1.Group("/company")
		{
			company.POST("/", r.Controller.CreateCompany)
		}
		user := v1.Group("/user")
		{
			user.POST("/", r.Controller.CreateUser)
		}

	}

	// Invalid URL
	r.Gin.NoRoute(func(c *gin.Context) {
		logs.WithFields(logs.Fields{"URL": r.address + r.port, "Method": c.Request.Method, "Path": c.Request.URL.Path}).Error("[NoRoute]Invalid")
		c.JSON(404, gin.H{"code": "404", "message": "Page not found"})
	})

	r.Logs.Info("initializing Router done")
}
