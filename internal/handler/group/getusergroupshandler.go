package group

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/ryantokmanmok/chat-app-server/common/errx"

	"net/http"

	"github.com/ryantokmanmok/chat-app-server/internal/logic/group"
	"github.com/ryantokmanmok/chat-app-server/internal/svc"
	"github.com/ryantokmanmok/chat-app-server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserGroupsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserGroupReq
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

		l := group.NewGetUserGroupsLogic(r.Context(), svcCtx)
		resp, err := l.GetUserGroups(&req)
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
