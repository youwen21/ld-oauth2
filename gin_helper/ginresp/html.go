package ginresp

import (
	"fmt"
)

func PageOk(timeout int, uri string, title string, msg string) string {
	//c.String(http.StatusOK, pageRefresh, timeout, uri, title, msg, timeout)
	pageHtml := fmt.Sprintf(pageRefresh, timeout, uri, title, msg, timeout)
	return pageHtml
	//c.Write()
}

func PageErr(timeout int, uri string, title string, msg string) string {
	//c.String(http.StatusOK, pageRefresh, timeout, uri, title, msg, timeout)
	pageHtml := fmt.Sprintf(pageRefresh, timeout, uri, title, msg, timeout)
	return pageHtml
	//c.Write()
}

func PageRefresh(timeout int, uri string, title string, msg string) string {
	//c.String(http.StatusOK, pageRefresh, timeout, uri, title, msg, timeout)
	pageHtml := fmt.Sprintf(pageRefresh, timeout, uri, title, msg, timeout)
	return pageHtml
	//c.Write()
}
