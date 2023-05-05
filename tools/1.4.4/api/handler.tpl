package {{.PkgName}}

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
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
        en_translations.RegisterDefaultTranslations(validate, trans)

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
            if e, ok := err.(*errx.CustomError); ok {
               httpx.WriteJsonCtx(r.Context(), w, e.StatusCode(), e.ToJSON())
            } else {
                httpx.ErrorCtx(r.Context(), w, err)
            }
		} else {
			{{if .HasResp}}httpx.OkJsonCtx(r.Context(), w, resp){{else}}httpx.Ok(w){{end}}
		}
	}
}
