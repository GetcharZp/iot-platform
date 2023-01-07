package svc

import (
	"gitee/getcharzp/iot-platform/device/device"
	"gitee/getcharzp/iot-platform/open/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	RpcDevice device.Device
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		RpcDevice: device.NewDevice(zrpc.MustNewClient(c.RpcClientConf)),
	}
}
