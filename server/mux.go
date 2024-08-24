package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"gofly/asset"
	"gofly/router"
	"io/fs"
	"net/http"
)

var (
	Mux      *http.ServeMux
	MyServer server
)

type server struct {
	distServer   FsServerI
	staticServer FsServerI

	ginEngine *gin.Engine
}

func (f server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// staticServer 和 distServer 的CanServe 有 /的逻辑
	// 如果这两个静态服务都不支持 /
	// 走本函数中的 兜底 访问首页的逻辑

	if f.staticServer.CanServe(r.URL.Path) {
		f.staticServer.ServeHTTP(w, r)
		return
	}

	if f.distServer.CanServe(r.URL.Path) {
		f.distServer.ServeHTTP(w, r)
		return
	}

	// 兜底 访问首页的逻辑
	if r.URL.Path == "/" {
		newR := r.Clone(context.Background())
		newR.URL.Path = "/default.html"
		f.staticServer.ServeHTTP(w, newR)
		//http.Redirect(w, r, "/square.html", http.StatusFound)
		return
	}

	// gin api
	f.ginEngine.ServeHTTP(w, r)
}

func HandleFavicon(w http.ResponseWriter, r *http.Request) {
	cnt, err := asset.Favicon.ReadFile("favicon_io/favicon.ico")
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "image/x-icon")
	w.Header().Set("Cache-Control", "max-age=84600, public")
	w.Write(cnt)
}

func init() {
	distFS, _ := fs.Sub(asset.Dist, "dist")
	staticFS, _ := fs.Sub(asset.Static, "static")

	ginEngine := gin.Default()
	router.InitRouter(ginEngine)

	MyServer = server{
		distServer:   NewLocalOrEmbedServer("./asset/dist/", distFS),
		staticServer: NewLocalOrEmbedServer("./asset/static/", staticFS),
		ginEngine:    ginEngine,
	}

	Mux = http.NewServeMux()
	Mux.Handle("/", MyServer)
	Mux.HandleFunc("/favicon.ico", HandleFavicon)
	//mux.Handle("/xxx", otherServer) 在myServer之上再封状一层
}
