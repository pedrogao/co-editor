//go:build js && wasm
// +build js,wasm

package main

import (
	"log"
	"net/url"
	"syscall/js"

	"ot"
)

func main() {
	window := js.Global()
	document := window.Get("document")
	body := document.Get("body")
	location := window.Get("location").String()
	loc, err := url.Parse(location)
	if err != nil {
		log.Fatal(err)
	}
	wsLoc, err := url.Parse(loc.String())
	wsLoc.Scheme = "ws"
	wsLoc.Path = "/ws"
	ot.NewTextarea(document, body, wsLoc.String(), "crhntr")
	select {}
}
