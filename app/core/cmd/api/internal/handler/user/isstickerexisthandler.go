package user

import (
	"net/http"

	"api/app/core/cmd/api/internal/logic/user"
	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// Check an existing sticker has been added to user
func IsStickerExistHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IsStickerExistReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewIsStickerExistLogic(r.Context(), svcCtx)
		resp, err := l.IsStickerExist(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}