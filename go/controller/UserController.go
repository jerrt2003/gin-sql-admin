package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jerrt2003/go/gin-sql-admin/go/entity"
	"github.com/jerrt2003/go/gin-sql-admin/go/service"
)

func CreateUser(c *gin.Context) {
	var user entity.User
	c.BindJSON(&user)
	err := service.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": user,
		})
	}
}

func GetUserList(c *gin.Context) {
	userList, err := service.GetAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": userList,
		})
	}
}
