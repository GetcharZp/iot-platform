package svc

import (
	"gitee/getcharzp/iot-platform/device/internal/config"
	"gitee/getcharzp/iot-platform/models"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB()
	return &ServiceContext{
		Config: c,
		DB:     models.DB,
	}
}
