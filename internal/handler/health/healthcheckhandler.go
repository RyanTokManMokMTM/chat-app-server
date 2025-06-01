package health

import (
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"

	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/logic/health"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HealthCheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := health.NewHealthCheckLogic(r.Context(), svcCtx)
		resp, err := l.HealthCheck()
		if err != nil {
			//convert to customError
			if e, ok := err.(*errx.CustomError); ok {
				httpx.WriteJsonCtx(r.Context(), w, e.StatusCode(), e.ToJSON())
			} else {
				httpx.ErrorCtx(r.Context(), w, err)
			}
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
