package middlewares

import (
	"net/http"

	"github.com/faruqisan/social-info/auth"
)

// CheckGoogleAccessToken .
func CheckGoogleAccessToken(googleAPI auth.Auth, next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !googleAPI.CheckToken() {
			http.Redirect(w, r, googleAPI.GetAuthorizeURL(), http.StatusFound)
		} else {
			next.ServeHTTP(w, r)
		}

	})

}
