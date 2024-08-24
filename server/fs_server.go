package server

import (
	"bufio"
	"io"
	"io/fs"
	"net/http"
	"regexp"
	"strings"
)

//http.StripPrefix("/admin/", http.FileServer(http.FS(sub)))

type FsServer struct {
	Fs fs.FS

	HistoryRouters HistoryRouters

	fileServer http.Handler
}

func (s FsServer) CanServe(urlPath string) bool {
	if s.HistoryRouters != nil && s.HistoryRouters.IsContain(urlPath) {
		return true
	}

	upath := strings.TrimLeft(urlPath, "/")
	// 二.2  static embed
	_, err := s.Fs.Open(upath)
	if err == nil {
		return true
	}

	return false
}

func (s FsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 非文件，而是vue路由， 输出index.html首页
	if s.HistoryRouters != nil && s.HistoryRouters.IsContain(r.URL.Path) {
		// 文件不存在 输出首页
		fi, _ := s.Fs.Open("index.html")
		defer fi.Close()
		content, _ := io.ReadAll(fi)
		w.Write(content)
		return
	}

	// css, js, img等文件
	// 需要输出正确的header头，所以借助http.FileServer
	s.fileServer.ServeHTTP(w, r)
}

type FsServerI interface {
	CanServe(urlPath string) bool
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type HistoryRouters []string

func (p HistoryRouters) IsContain(path string) bool {
	for _, v := range p {
		if strings.Contains(v, "\\") {
			ok, err := regexp.MatchString(v, path)
			if ok && err == nil {
				return true
			}
		}

		if v == path {
			return true
		}
	}
	return false
}

func initHistoryRouters(file fs.File) (HistoryRouters, error) {
	hsRouters := make(HistoryRouters, 0)
	r := bufio.NewReader(file)
	for {
		// ReadLine is a low-level line-reading primitive.
		// Most callers should use ReadBytes('\n') or ReadString('\n') instead or use a Scanner.
		bytes, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		hsRouters = append(hsRouters, string(bytes))
	}

	return hsRouters, nil
}
