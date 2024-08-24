package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	methods = `POST, OPTIONS, GET, PUT, DELETE`
	headers = `Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization,Accept,Origin,Men,Cache-Control,X-Requested-With,Name,DNT,HOST,Pragma,Referer,Duo,Range,user-Agent,token`
)

type cors struct {
}

var Cors = &cors{}

func (c *cors) RawCors(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("X-ORIGIN")
		if origin == "" {
			origin = r.Header.Get("ORIGIN")
		}

		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
		w.Header().Set("Access-Control-Allow-Methods", methods)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", headers)

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (c *cors) GinCors() gin.HandlerFunc {
	// CORSMiddleware 跨域
	// @see https://stackoverflow.com/questions/29418478/go-gin-framework-cors
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("X-ORIGIN")
		if origin == "" {
			origin = c.Request.Header.Get("ORIGIN")
		}

		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", methods)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", headers)

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
