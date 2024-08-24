package server

import (
	"errors"
	"io/fs"
	"net/http"
	"os"
)

func NewEmbedServer(distFs fs.FS) FsServerI {
	if distFs == nil {
		panic(errors.New("init NewEmbedServer fail"))
	}

	var routers HistoryRouters
	f, err := distFs.Open("httphere_routers.txt")
	if err == nil {
		defer f.Close()
		routers, _ = initHistoryRouters(f)
	}

	// 不存在本地 ./asset/static/目录
	return &FsServer{
		distFs,
		routers,
		http.FileServer(http.FS(distFs)),
	}

}

func NewLocalServer(localDir string) FsServerI {
	_, err := os.Stat(localDir)
	if err != nil {
		panic(errors.New("init NewLocalServer fail:" + localDir))
	}

	localFs := os.DirFS(localDir)
	var routers HistoryRouters
	f, err := localFs.Open("httphere_routers.txt")
	if err == nil {
		defer f.Close()
		routers, _ = initHistoryRouters(f)
	}

	return &FsServer{
		os.DirFS(localDir),
		routers,
		http.FileServer(http.Dir(localDir)),
	}
}

func NewLocalOrEmbedServer(localDir string, distFs fs.FS) FsServerI {
	if localDir != "" {
		if _, err := os.Stat(localDir); err == nil {
			return NewLocalServer(localDir)
		}
	}

	if distFs != nil {
		return NewEmbedServer(distFs)
	}

	panic(errors.New("init NewLocalOrEmbedServer fail"))
}
