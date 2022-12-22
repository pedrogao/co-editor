package handlers

import (
	"encoding/json"
	"net/http"

	"woot/woot"

	log "github.com/sirupsen/logrus"
)

func SyncDelete(w http.ResponseWriter, r *http.Request) {
	var err error
	var requestBody deleteBody
	var wStringPointer *woot.WString

	// Obtain the value & position from POST Request Body
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&requestBody)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed parse request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// SyncDelete the given value in our stored WString
	wStringPointer = WString.GenerateDelete(requestBody.Position)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to delete sync value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	WString = *wStringPointer

	log.WithFields(log.Fields{
		"text":     woot.Value(WString),
		"position": requestBody.Position,
	}).Debug("successful wstring sync deletion")

	w.WriteHeader(http.StatusOK)
}
