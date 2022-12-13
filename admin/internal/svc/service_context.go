package svc

import (
	"gitee/getcharzp/iot-platform/admin/internal/config"
	"gitee/getcharzp/iot-platform/models"
	user2 "gitee/getcharzp/iot-platform/user/rpc/types/user"
	"gitee/getcharzp/iot-platform/user/rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gorm.DB
	RpcUser  user.User
	AuthUser *user2.UserAuthReply
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB()
	return &ServiceContext{
		Config:  c,
		DB:      models.DB,
		RpcUser: user.NewUser(zrpc.MustNewClient(c.RpcClientConf)),
	}
}
