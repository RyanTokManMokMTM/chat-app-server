package {{.PkgName}}

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
    "github.com/pkg/errors"
    "google.golang.org/grpc/status"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	"net/http"

	f"github.com/zeromicro/go-zero/rest/httpx"
	{{.ImportPackages}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
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


		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
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
			{{if .HasResp}}httpx.OkJsonCtx(r.Context(), w, resp){{else}}httpx.Ok(w){{end}}
		}
	}
}
