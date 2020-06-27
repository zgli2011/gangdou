package dao

type UserDao struct {
	DBEngine DBEngine
}

// AddUser: 添加用户
func (userDao *UserDao) AddUser(user *AuthUser) err {
	_, err := userDao.DBEngine.Insert(user)
	return err
}
