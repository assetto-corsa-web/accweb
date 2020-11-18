package pages

import (
	"github.com/assetto-corsa-web/accweb/auth"
	"github.com/assetto-corsa-web/accweb/server"
	"html/template"
	"log"
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
					Str: "sub",
					SubTestSlice: []server.SubTest{
						{
							Str: "sub sub",
							SubTestSlice: []server.SubTest{
								{
									Str: "sub sub sub",
								},
							},
						},
						{},
						{},
					},
				},
			},
		}),
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			// TODO
			panic(err)
		}

		values := r.Form
		log.Println(values.Get("udpPort"))
		log.Println(values.Get("select"))

		// TODO redirect to overview
	}

	if err := executeTemplate(w, "server.html", data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
