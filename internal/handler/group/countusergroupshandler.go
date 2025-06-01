package group

import (
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/logic/group"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// Count user group
func CountUserGroupsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CountUserGroupsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := group.NewCountUserGroupsLogic(r.Context(), svcCtx)
		resp, err := l.CountUserGroups(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
