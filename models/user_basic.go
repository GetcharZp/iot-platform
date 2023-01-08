package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity  string `gorm:"column:identity; type:varchar(50);" json:"identity"`
	Name      string `gorm:"column:name; type:varchar(50);" json:"name"`
	Password  string `gorm:"column:password; type:varchar(50);" json:"password"`
	AppKey    string `gorm:"column:app_key; type:varchar(36);" json:"app_key"`
	AppSecret string `gorm:"column:app_secret; type:varchar(36);" json:"app_secret"`
}

func (table UserBasic) TableName() string {
	return "user_basic"
}
