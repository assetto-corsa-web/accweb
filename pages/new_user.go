package pages

import (
	"github.com/assetto-corsa-web/accweb/auth"
	"net/http"
	"strings"
	"unicode"
)

type UserPageData struct {
	User          *auth.User
	Username      string
	Role          string
	UsernameError string
	PasswordError string
	RoleError     string
}

func NewUser(w http.ResponseWriter, r *http.Request, claims *auth.TokenClaims) {
	if !claims.IsAdmin {
		http.Redirect(w, r, "/not-found", http.StatusNotFound)
		return
	}

	data := UserPageData{}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data.Username = strings.TrimSpace(r.Form.Get("username"))
		password := r.Form.Get("password")
		repeatPassword := r.Form.Get("repeat_password")
		data.Role = r.Form.Get("role")
		data.UsernameError = isValidUsername(data.Username)
		data.PasswordError = isValidPassword(password, repeatPassword)
		data.RoleError = isValidRole(data.Role)

		if data.UsernameError == "" && data.PasswordError == "" && data.RoleError == "" {
			auth.GetUserList().Set(data.Username, password, data.Role)
			http.Redirect(w, r, "/user", http.StatusFound)
			return
		}
	}

	if err := executeTemplate(w, "new_user.html", data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func isValidUsername(username string) string {
	if username == "" {
		return "Username must not be empty."
	}

	for _, c := range []rune(username) {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != ' ' && c != '-' && c != '_' {
			return "Username must only consist out of letters, numbers, or spaces."
		}
	}

	return ""
}

func isValidPassword(password, repeatPassword string) string {
	if password == "" {
		return "Password must not be empty."
	}

	if password != repeatPassword {
		return "Passwords do not match."
	}

	return ""
}

func isValidRole(role string) string {
	if role == "admin" || role == "moderator" || role == "read-only" {
		return ""
	}

	return "Role invalid."
}
