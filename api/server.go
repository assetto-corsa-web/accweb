package api

import (
	"github.com/assetto-corsa-web/accweb/server"
	"net/http"
)

func SaveServerSetttingsHandler(w http.ResponseWriter, r *http.Request) {
	req := &server.ServerSettings{}

	if err := decodeJSON(r, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := server.SaveServerSettings(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	writeResponse(w, nil)
}
