package assets

import (
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/api/internal/logic/assets"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// Upload any file
func UploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadFileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := assets.NewUploadFileLogic(r.Context(), svcCtx, r)
		resp, err := l.UploadFile(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
