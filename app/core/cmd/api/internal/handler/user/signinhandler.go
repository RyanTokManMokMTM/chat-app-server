package user

import (
	"api/app/common/errx"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"net/http"

	"api/app/core/cmd/api/internal/logic/user"
	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// User account sign in
func SignInHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignInReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewSignInLogic(r.Context(), svcCtx)
		resp, err := l.SignIn(&req)
		if err != nil {
			//RPC error or custom error
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
			logx.WithContext(r.Context()).Errorf("User Sign In error : %+v", reqError)
			httpx.ErrorCtx(r.Context(), w, reqError)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
