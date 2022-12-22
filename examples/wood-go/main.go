package main

import (
	"net/http"
	"os"

	"woot/handlers"

	log "github.com/sirupsen/logrus"
)

const (
	PORT = "8080"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	r := handlers.Router()

	log.WithFields(log.Fields{
		"port": PORT,
	}).Info("started WOOT node server")

	http.ListenAndServe(":"+PORT, r)
}
