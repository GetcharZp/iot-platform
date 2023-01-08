package logic

import (
	"context"
	"encoding/json"
	"errors"
	"gitee/getcharzp/iot-platform/helper"
	"gitee/getcharzp/iot-platform/models"
	"sort"

	"gitee/getcharzp/iot-platform/user/rpc/internal/svc"
	"gitee/getcharzp/iot-platform/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type OpenAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOpenAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenAuthLogic {
	return &OpenAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OpenAuthLogic) OpenAuth(in *user.OpenAuthRequest) (*user.OpenAuthReply, error) {
	data := make(map[string]interface{})
	err := json.Unmarshal(in.Body, &data)
	if err != nil {
		logx.Error("[Unmarshal ERROR] : ", err.Error())
		return nil, err
	}

	ub := new(models.UserBasic)
	err = l.svcCtx.DB.Model(new(models.UserBasic)).Select("app_secret").Where("app_key = ?", data["app_key"]).First(ub).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err.Error())
		return nil, err
	}

	arr := make([]string, 0)
	for k, _ := range data {
		arr = append(arr, k)
	}
	sort.Strings(arr)

	var s string
	for _, v := range arr {
		if v != "sign" {
			s += data[v].(string)
		}
	}

	if helper.Md5(s) != data["sign"].(string) {
		return nil, errors.New("签名不正确")
	}

	return &user.OpenAuthReply{}, nil
}
