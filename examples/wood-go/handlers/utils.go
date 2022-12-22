package handlers

import (
	"os"
	"strings"
)

func GetPeerList() []string {
	if os.Getenv("PEERS") == "" {
		return []string{}
	}
	return strings.Split(os.Getenv("PEERS"), ",")
}

func GetNetwork() string {
	return os.Getenv("NETWORK") + ":8080"
}
