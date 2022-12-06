package models

import "gorm.io/gorm"

type DeviceBasic struct {
	gorm.Model
	Identity        string `gorm:"column:identity; type:varchar(50);" json:"identity"`
	ProductIdentity string `gorm:"column:product_identity; type:varchar(50);" json:"product_identity"`
	Name            string `gorm:"column:name; type:varchar(50);" json:"name"`
	Key             string `gorm:"column:key; type:varchar(50);" json:"key"`
	Secret          string `gorm:"column:secret; type:varchar(50);" json:"secret"`
}

func (table DeviceBasic) TableName() string {
	return "device_basic"
}
