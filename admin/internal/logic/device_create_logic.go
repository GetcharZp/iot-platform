package logic

import (
	"context"
	"gitee/getcharzp/iot-platform/admin/internal/svc"
	"gitee/getcharzp/iot-platform/admin/internal/types"
	"gitee/getcharzp/iot-platform/models"
	"github.com/google/uuid"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceCreateLogic {
	return &DeviceCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceCreateLogic) DeviceCreate(req *types.DeviceCreateRequest) (resp *types.DeviceCreateReply, err error) {
	err = l.svcCtx.DB.Create(&models.DeviceBasic{
		Identity:        uuid.New().String(),
		ProductIdentity: req.ProductIdentity,
		Name:            req.Name,
		Key:             uuid.New().String(),
		Secret:          uuid.New().String(),
	}).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
	}

	return
}
