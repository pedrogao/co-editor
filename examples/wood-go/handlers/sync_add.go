package handlers

import (
	"encoding/json"
	"net/http"

	"woot/woot"

	log "github.com/sirupsen/logrus"
)

func SyncAdd(w http.ResponseWriter, r *http.Request) {
	var err error
	var requestBody addBody
	var wStringPointer *woot.WString

	// Obtain the value & position from POST Request Body
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&requestBody)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed parse request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// SyncAdd the given value to our stored WString
	wStringPointer, err = WString.GenerateInsert(requestBody.Position, requestBody.Value)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to add sync value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	WString = *wStringPointer

	log.WithFields(log.Fields{
		"text":     woot.Value(WString),
		"value":    requestBody.Value,
		"position": requestBody.Position,
	}).Debug("successful wstring sync addition")

	w.WriteHeader(http.StatusOK)
}
