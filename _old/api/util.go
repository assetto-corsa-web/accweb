package api

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

func decodeJSON(r *http.Request, req interface{}) error {
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&req); err != nil {
		logrus.WithField("err", err).Debug("Error decoding JSON request")
		return errors.New("Invalid format")
	}

	logrus.WithField("req", req).Debug("Decoded JSON request")
	return nil
}

func writeResponse(w http.ResponseWriter, resp interface{}) {
	// send an empty response by default
	if resp == nil {
		resp = struct{}{}
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		logrus.WithFields(logrus.Fields{"err": err, "response": resp}).Error("Error marshalling response")
		w.WriteHeader(http.StatusInternalServerError)
	}
}
