package models

import "time"

// AuthUser : auth_user表映射集合
type AuthUser struct {
	ID          int
	Username    string    `xorm:"varchar(64)"`
	Password    string    `xorm:"varchar(255)"`
	UserType    int       `xorm:"int"`
	RealName    string    `xorm:"varchar(64)"`
	Mail        string    `xorm:"varchar(64)"`
	Mobile      string    `xorm:"varchar(16)"`
	Status      int       `xorm:"int"`
	Description string    `xorm:"varchar(255)"`
	Creator     string    `xorm:"varchar(64)"`
	CreateTime  time.Time `xorm:"created"`
	UpdateTime  time.Time `xorm:"updated"`
}
