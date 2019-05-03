package api

import (
	"github.com/assetto-corsa-web/accweb/server"
	"net/http"
	"strconv"
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

func GetServerListHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		writeResponse(w, server.GetServerList())
	} else {
		idInt, err := strconv.Atoi(id)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		writeResponse(w, server.GetServerById(idInt))
	}
}
