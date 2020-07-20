package auth

import (
	"github.com/emvi/logbuch"
	"net/http"
)

const (
	TokenCookie = "token"
)

type HttpHandler func(http.ResponseWriter, *http.Request, *TokenClaims)

// Middleware reads the JWT from the request cookie
// and continues serving the page in case it is valid or redirects to login in case it is not.
func Middleware(next HttpHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := getTokenFromCookie(r)
		claims := isValidToken(token)

		if claims == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		next(w, r, claims)
	})
}

func getTokenFromCookie(r *http.Request) string {
	cookie, err := r.Cookie(TokenCookie)

	if err != nil {
		logbuch.Debug("Error reading token cooke", logbuch.Fields{"err": err})
		return ""
	}

	return cookie.Value
}
