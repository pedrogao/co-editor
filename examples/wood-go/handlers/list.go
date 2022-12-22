package handlers

import (
	"encoding/json"
	"net/http"

	"woot/woot"

	log "github.com/sirupsen/logrus"
)

func List(w http.ResponseWriter, r *http.Request) {
	text := woot.Value(WString)

	log.WithFields(log.Fields{
		"text": text,
	}).Debug("successful wstring list")

	json.NewEncoder(w).Encode(text)
}
