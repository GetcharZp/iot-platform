package handler

import (
	"net/http"

	"gitee/getcharzp/iot-platform/admin/internal/logic"
	"gitee/getcharzp/iot-platform/admin/internal/svc"
	"gitee/getcharzp/iot-platform/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ProductDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductDeleteRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewProductDeleteLogic(r.Context(), svcCtx)
		resp, err := l.ProductDelete(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
