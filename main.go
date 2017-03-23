package main

import (
	"log"
	"net/http"
)

func main() {
	startWebSocket()
	startHTTP()
	if err := http.ListenAndServe(config.HTTPAddress, nil); err != nil {
		log.Fatal(err)
	}
}
