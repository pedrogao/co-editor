package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"woot/woot"
)

var (
	WString woot.WString
)

func init() {
	WString = woot.Initialize()
}

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

var Routes = []Route{
	{"/woot", "GET", WOOT},
	{"/woot/list", "GET", List},
	{"/woot/add", "POST", Add},
	{"/woot/delete", "POST", Delete},
	{"/woot/sync/add", "POST", SyncAdd},
	{"/woot/sync/delete", "POST", SyncDelete},
}

func WOOT(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World WOOT Node\n")
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"path":   r.URL,
			"method": r.Method,
		}).Info("incoming request")
		next.ServeHTTP(w, r)
	})
}

func Router() *mux.Router {
	router := mux.NewRouter()

	for _, route := range Routes {
		router.HandleFunc(
			route.Path,
			route.Handler,
		).Methods(route.Method)
	}

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./handlers/public")))
	router.Use(Logger)

	return router
}
