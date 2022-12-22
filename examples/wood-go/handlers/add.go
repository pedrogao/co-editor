package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"woot/woot"
)

type addBody struct {
	Value    string `json:"value"`
	Position int    `json:"position"`
}

func Add(w http.ResponseWriter, r *http.Request) {
	var (
		err            error
		requestBody    addBody
		wStringPointer *woot.WString
	)

	decoder := json.NewDecoder(r.Body)
	if err = decoder.Decode(&requestBody); err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to add value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	wStringPointer, err = WString.GenerateInsert(requestBody.Position, requestBody.Value)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to add value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	WString = *wStringPointer

	log.WithFields(log.Fields{
		"text":     woot.Value(WString),
		"value":    requestBody.Value,
		"position": requestBody.Position,
	}).Debug("successful wstring addition")

	err = BroadcastAdd(requestBody.Position, requestBody.Value)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to sync add value")
	}

	w.WriteHeader(http.StatusOK)
}
