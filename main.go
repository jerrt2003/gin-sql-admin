package main

import (
	"os"

	"github.com/jerrt2003/gin-sql-admin/go/dao"
	"github.com/jerrt2003/gin-sql-admin/go/entity"
	"github.com/jerrt2003/gin-sql-admin/go/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	args := os.Args[1:]
	sqlConfig := args[0]
	err := dao.InitMySql(sqlConfig)
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	// 為什麼這邊建的dao.SqlSession其他的也可以用呢:(我的想法)因為在Go同一個pkg下都是所有的func都可以被main access
	dao.SqlSession.AutoMigrate(&entity.User{})
	r := routes.SetRouter()
	r.Run(":8081")
}
