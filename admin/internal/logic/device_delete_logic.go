package logic

import (
	"context"
	"gitee/getcharzp/iot-platform/api"
	"gitee/getcharzp/iot-platform/models"
	"gorm.io/gorm"

	"gitee/getcharzp/iot-platform/admin/internal/svc"
	"gitee/getcharzp/iot-platform/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceDeleteLogic {
	return &DeviceDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceDeleteLogic) DeviceDelete(req *types.DeviceDeleteRequest) (resp *types.DeviceDeleteReply, err error) {
	deviceBasic := new(models.DeviceBasic)
	err = l.svcCtx.DB.Model(new(models.DeviceBasic)).Select("key").Where("identity = ?", req.Identity).
		Find(deviceBasic).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return
	}

	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 数据库中删除设备
		err = tx.Where("identity = ?", req.Identity).Delete(new(models.DeviceBasic)).Error
		if err != nil {
			logx.Error("[DB ERROR] : ", err)
			return err
		}

		// 2. EMQX 中同步删除认证设备
		err = api.DeleteAuthUser(deviceBasic.Key)
		if err != nil {
			logx.Error("[DeleteAuthUser ERROR] : ", err)
			return err
		}
		return nil
	})

	return
}
