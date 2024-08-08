package assets

import (
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/api/internal/logic/assets"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// Upload only image
func UploadImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadImageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := assets.NewUploadImageLogic(r.Context(), svcCtx)
		resp, err := l.UploadImage(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
