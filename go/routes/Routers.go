package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jerrt2003/go/gin-sql-admin/go/controller"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("user")
	userGroup.POST("/users", controller.CreateUser)
	userGroup.GET("/users", controller.GetUserList)

	return r
}
