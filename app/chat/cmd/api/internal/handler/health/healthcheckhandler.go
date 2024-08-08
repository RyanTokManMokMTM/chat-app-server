package health

import (
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"google.golang.org/grpc/status"

	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/app/chat/cmd/api/internal/logic/health"
	"github.com/ryantokmanmokmtm/chat-app-server/app/chat/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HealthCheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := health.NewHealthCheckLogic(r.Context(), svcCtx)
		resp, err := l.HealthCheck()
		if err != nil {
			//convert to customError
			reqError := errx.NewCustomErrCode(errx.SERVER_COMMON_ERROR)
			errType := errors.Cause(err)
			var e *errx.CustomError
			if errors.As(errType, &e) {
				//Custom error
				reqError = e
			} else {
				//gRPC error
				if statusRPC, ok := status.FromError(errType); ok {
					grpcCode := errx.InternalCode(statusRPC.Code())
					//Is defined error or underlay error?
					if errx.IsErrorCode(grpcCode) {
						reqError = errx.NewCustomErrCode(grpcCode)
					}
				}
			}
			httpx.WriteJsonCtx(r.Context(), w, reqError.StatusCode(), reqError.ToJSON())
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
