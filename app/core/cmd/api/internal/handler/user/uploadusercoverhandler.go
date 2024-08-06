package user

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"google.golang.org/grpc/status"

	"net/http"

	"api/app/core/cmd/api/internal/logic/user"
	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadUserCoverHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadUserAvatarReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		eng := en.New()
		uti := ut.New(eng, eng)
		trans, _ := uti.GetTranslator("en")
		validate := validator.New()
		err := en_translations.RegisterDefaultTranslations(validate, trans)
		if err != nil {
			commonErr := errx.NewCustomErrCode(errx.SERVER_COMMON_ERROR)
			httpx.WriteJsonCtx(r.Context(), w, commonErr.StatusCode(), commonErr.ToJSON())
			return
		}
		if err := validate.StructCtx(r.Context(), req); err != nil {
			err := err.(validator.ValidationErrors)
			commonErr := errx.NewCustomError(errx.REQ_PARAM_ERROR, err[0].Translate(trans))
			httpx.WriteJsonCtx(r.Context(), w, commonErr.StatusCode(), commonErr.ToJSON())
			return
		}

		l := user.NewUploadUserCoverLogic(r.Context(), svcCtx, r)
		resp, err := l.UploadUserCover(&req)
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
