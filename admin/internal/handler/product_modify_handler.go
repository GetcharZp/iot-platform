package handler

import (
	"net/http"

	"gitee/getcharzp/iot-platform/admin/internal/logic"
	"gitee/getcharzp/iot-platform/admin/internal/svc"
	"gitee/getcharzp/iot-platform/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ProductModifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductModifyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewProductModifyLogic(r.Context(), svcCtx)
		resp, err := l.ProductModify(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
