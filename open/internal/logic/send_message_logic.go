package logic

import (
	"context"
	"gitee/getcharzp/iot-platform/device/types/device"

	"gitee/getcharzp/iot-platform/open/internal/svc"
	"gitee/getcharzp/iot-platform/open/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendMessageLogic) SendMessage(req *types.SendMessageRequest) (resp *types.SendMessageReply, err error) {
	_, err = l.svcCtx.RpcDevice.SendMessage(l.ctx, &device.SendMessageRequest{
		ProductKey: req.ProductKey,
		DeviceKey:  req.DeviceKey,
		Data:       req.Data,
	})
	if err != nil {
		logx.Error("[ERROR] : ", err.Error())
	}
	return
}
