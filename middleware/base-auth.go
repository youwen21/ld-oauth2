package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// What is the "realm" in basic authentication
// https://stackoverflow.com/questions/12701085/what-is-the-realm-in-basic-authentication

//How to Implement HTTP Basic Auth in Go(Golang)
// https://umesh.dev/blog/how-to-implement-http-basic-auth-in-gogolang/

// How to correctly use Basic Authentication in Go
// https://www.alexedwards.net/blog/basic-authentication-in-go
type baseAuth struct {
	Realm string

	Username string
	Password string
}

var BaseAuth = &baseAuth{Realm: "member", Username: "admin", Password: "admin123456!!"}

func (auth *baseAuth) BasicAuth(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte(auth.Username))
			expectedPasswordHash := sha256.Sum256([]byte(auth.Password))

			usernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1
			passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s", charset="UTF-8"`, auth.Realm))
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

func (auth *baseAuth) GinBasicAuth(c *gin.Context) {
	username, password, ok := c.Request.BasicAuth()
	if ok {
		usernameHash := sha256.Sum256([]byte(username))
		passwordHash := sha256.Sum256([]byte(password))
		expectedUsernameHash := sha256.Sum256([]byte(auth.Username))
		expectedPasswordHash := sha256.Sum256([]byte(auth.Password))

		usernameMatch := subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1
		passwordMatch := subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1

		if usernameMatch && passwordMatch {
			c.Next()
			return
		}
	}

	c.Writer.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s", charset="UTF-8"`, auth.Realm))
	http.Error(c.Writer, "Unauthorized", http.StatusUnauthorized)
	c.Abort()
}
