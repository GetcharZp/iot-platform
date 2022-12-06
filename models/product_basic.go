package models

import "gorm.io/gorm"

type ProductBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity; type:varchar(50);" json:"identity"`
	Name     string `gorm:"column:name; type:varchar(50);" json:"name"`
	Desc     string `gorm:"column:desc; type:varchar(50);" json:"desc"`
}

func (table ProductBasic) TableName() string {
	return "product_name"
}
