package logic

import (
	"context"

	"gitee/getcharzp/iot-platform/device/internal/svc"
	"gitee/getcharzp/iot-platform/device/types/device"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMessageLogic) SendMessage(in *device.SendMessageRequest) (*device.SendMessageReply, error) {
	// todo: add your logic here and delete this line

	return &device.SendMessageReply{}, nil
}
