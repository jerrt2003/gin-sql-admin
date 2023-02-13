package dao

import (
	"fmt"
	"io/ioutil"

	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
)

const DRIVER = "mysql"

var SqlSession *gorm.DB

type conf struct {
	Url      string `yaml: "url"`
	UserName string `yaml: "username"`
	Password string `yaml: "password"`
	DbName   string `yaml: "dbname"`
	Port     string `yaml: "port`
}

func (c *conf) getConf(filePath string) *conf {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

func InitMySql(filePath string) (err error) {
	var c conf
	conf := c.getConf(filePath)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName,
		conf.Password,
		conf.Url,
		conf.Port,
		conf.DbName,
	)
	SqlSession, err = gorm.Open(DRIVER, dsn)
	if err != nil {
		panic(err)
	}
	return SqlSession.DB().Ping()
}

func Close() {
	SqlSession.Close()
}
