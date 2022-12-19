//go:build !js || !wasm
// +build !js !wasm

package ot

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

type Server struct {
	Listeners         []chan []Applier
	RegisterListeners chan chan []Applier
}

func (server *Server) HandlerWS(res http.ResponseWriter, req *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(req, res)
	if err != nil {
		log.Println("could not upgreade connection")
		log.Print(err)
		return
	}
	defer conn.Close()
	for {
		msg, op, err := wsutil.ReadClientData(conn)
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Println(string(msg))
		err = wsutil.WriteServerMessage(conn, op, msg)
		if err != nil {
			log.Println(err)
			break
		}
	}
}
