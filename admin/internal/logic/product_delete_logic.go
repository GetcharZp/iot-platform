package logic

import (
	"context"
	"errors"
	"gitee/getcharzp/iot-platform/models"

	"gitee/getcharzp/iot-platform/admin/internal/svc"
	"gitee/getcharzp/iot-platform/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDeleteLogic {
	return &ProductDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductDeleteLogic) ProductDelete(req *types.ProductDeleteRequest) (resp *types.ProductDeleteReply, err error) {
	var cnt int64
	err = l.svcCtx.DB.Model(new(models.DeviceBasic)).Where("product_identity = ?", req.Identity).
		Count(&cnt).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return
	}
	if cnt > 0 {
		err = errors.New("已关联设备，不可删除")
		return
	}

	err = l.svcCtx.DB.Debug().Where("identity = ?", req.Identity).Delete(new(models.ProductBasic)).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
	}
	return
}
