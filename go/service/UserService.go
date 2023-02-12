package service

import (
	"github.com/jerrt2003/gin-sql-admin/go/dao"
	"github.com/jerrt2003/gin-sql-admin/go/entity"
)

func CreateUser(user *entity.User) (err error) {
	err = dao.SqlSession.Create(user).Error
	return
}

func GetAllUser() (userList []*entity.User, err error) {
	if err = dao.SqlSession.Find(&userList).Error; err != nil {
		return nil, err
	}
	return
}

func DeleteUserById(id string) (err error) {
	err = dao.SqlSession.Where("id=?", id).Delete(&entity.User{}).Error
	return
}

func GetUserById(id string) (user *entity.User, err error) {
	if err = dao.SqlSession.Where("id=?", id).First(user).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateUser(user *entity.User) (err error) {
	err = dao.SqlSession.Save(user).Error
	return
}
