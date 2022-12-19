package logic

import (
	"context"
	"gitee/getcharzp/iot-platform/admin/internal/svc"
	"gitee/getcharzp/iot-platform/admin/internal/types"
	"gitee/getcharzp/iot-platform/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductModifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductModifyLogic {
	return &ProductModifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductModifyLogic) ProductModify(req *types.ProductModifyRequest) (resp *types.ProductModifyReply, err error) {
	err = l.svcCtx.DB.Where("identity = ?", req.Identity).Updates(&models.ProductBasic{
		Name: req.Name,
		Desc: req.Desc,
	}).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
	}
	return
}
