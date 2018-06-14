package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/faruqisan/social-info/auth"
)

// CheckGoogleAccessToken middleware for check if google token already initialized
func CheckGoogleAccessToken(googleAPI auth.Auth) gin.HandlerFunc {

	return func(c *gin.Context) {
		if !googleAPI.CheckToken() {
			c.Redirect(http.StatusFound, googleAPI.GetAuthorizeURL())
		} else {
			c.Next()
		}
	}

}

// CheckGoogleAccessToken .
// func CheckGoogleAccessToken(googleAPI auth.Auth, next http.HandlerFunc) http.HandlerFunc {

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		if !googleAPI.CheckToken() {
// 			http.Redirect(w, r, googleAPI.GetAuthorizeURL(), http.StatusFound)
// 		} else {
// 			next.ServeHTTP(w, r)
// 		}

// 	})

// }
