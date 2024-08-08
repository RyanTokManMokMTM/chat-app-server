package router

import (
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/api/internal/config"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/api/internal/handler"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/util"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"path"
	"strings"
)

func ConfigRouter(c config.Config) *rest.Server {
	server := rest.MustNewServer(c.RestConf) // new a configRouter
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	dirPath := path.Join(util.GetRootDir(), "/resources/")
	dirRouterRegister(server, "/resources/", dirPath)

	return server
}

func directoryHandler(pattern, fileDir string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		h := http.StripPrefix(pattern, http.FileServer(http.Dir(fileDir)))
		h.ServeHTTP(w, req)

	}
}

func dirRouterRegister(server *rest.Server, pattern, dirPath string) {
	totalLevel := []string{
		":l1", ":l2", ":l3", ":l2", ":l4", ":l5", ":l6", ":l7", ":l8",
	}
	for i := 1; i < len(totalLevel); i++ {
		p := pattern + strings.Join(totalLevel[:i], "/")
		server.AddRoute(rest.Route{
			Method:  http.MethodGet,
			Path:    p,
			Handler: directoryHandler(pattern, dirPath),
		})
	}
}
