package svc

import (
	"gitee/getcharzp/iot-platform/models"
	"gitee/getcharzp/iot-platform/user/rpc/internal/config"
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
