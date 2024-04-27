package api

import (
	"playground/request"
	"playground/service"
	"playground/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func StartServer() {
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("valid-role", request.ValidRole)
	}
	authRoutes := router.Group("/").Use(authMiddleware())

	router.POST("/user", service.CreateUser)
	router.POST("/login", service.Login)

	authRoutes.GET("/user/:id", service.GetUser)
	authRoutes.GET("/user/", service.GetUserLists)

	router.Run(util.Config.ServerAddress)
}
