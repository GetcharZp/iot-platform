package handler

import (
	"net/http"
	"strconv"

	"gitee/getcharzp/iot-platform/admin/internal/logic"
	"gitee/getcharzp/iot-platform/admin/internal/svc"
	"gitee/getcharzp/iot-platform/admin/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeviceListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeviceListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		req.Page, _ = strconv.Atoi(r.URL.Query().Get("page"))
		req.Size, _ = strconv.Atoi(r.URL.Query().Get("size"))
		req.Name = r.URL.Query().Get("name")
		l := logic.NewDeviceListLogic(r.Context(), svcCtx)
		resp, err := l.DeviceList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
