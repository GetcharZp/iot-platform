package logic

import (
	"context"
	"errors"
	"gitee/getcharzp/iot-platform/helper"

	"gitee/getcharzp/iot-platform/user/rpc/internal/svc"
	"gitee/getcharzp/iot-platform/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthLogic) Auth(in *user.UserAuthRequest) (*user.UserAuthReply, error) {
	if in.Token == "" {
		return nil, errors.New("必填参不能为空")
	}
	userClaim, err := helper.AnalyzeToken(in.Token)
	if err != nil {
		return nil, err
	}
	resp := new(user.UserAuthReply)
	resp.Identity = userClaim.Identity
	resp.Id = uint64(userClaim.Id)
	resp.Extend = map[string]string{
		"name": userClaim.Name,
	}
	return resp, nil
}
