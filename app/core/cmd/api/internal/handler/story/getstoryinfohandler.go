package story

import (
	"net/http"

	"api/app/core/cmd/api/internal/logic/story"
	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// Get an instance story by storyID
func GetStoryInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetStoryInfoByIdRep
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := story.NewGetStoryInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetStoryInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
