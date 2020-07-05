package dao

import (
	"gangdou/models"

	"github.com/go-xorm/xorm"
)

// UserDao :
type UserDao struct {
	engine *xorm.Engine
}

// GetUsers : 添加用户
func (userDao *UserDao) GetUsers(id int64) []models.AuthUser {
	usersList := make([]models.AuthUser, 0)
	userDao.engine.Where("id=?", 1).Find(&usersList)
	return usersList
}
