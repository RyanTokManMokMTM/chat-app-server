package group

import (
	"net/http"

	"api/app/core/cmd/api/internal/logic/group"
	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// Update and update group info
func UpdateGroupInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateGroupInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := group.NewUpdateGroupInfoLogic(r.Context(), svcCtx)
		resp, err := l.UpdateGroupInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}