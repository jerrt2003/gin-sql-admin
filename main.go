package main

import (
	"github.com/jerrt2003/go/gin-sql-admin/go/dao"
	"github.com/jerrt2003/go/gin-sql-admin/go/entity"
	"github.com/jerrt2003/go/gin-sql-admin/go/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	err := dao.InitMySql()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	dao.SqlSession.AutoMigrate(&entity.User{})
	r := routes.SetRouter()
	r.Run(":8081")
}
