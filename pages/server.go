package pages

import (
	"github.com/assetto-corsa-web/accweb/auth"
	"github.com/assetto-corsa-web/accweb/server"
	"html/template"
	"net/http"
)

func Server(w http.ResponseWriter, r *http.Request, claims *auth.TokenClaims) {
	data := struct {
		IsAdmin bool
		Form    template.HTML
	}{
		claims.IsAdmin,
		server.Form(&server.Server{
			ConfigurationJSON: server.ConfigurationJSON{
				SubTest: server.SubTest{
					SubTestSlice: []server.SubTest{
						{},
						{},
						{},
					},
				},
			},
		}),
	}

	if err := executeTemplate(w, "server_page.html", data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
