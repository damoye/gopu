package main

import (
	"log"
	"net/http"

	"github.com/damoye/gopu/bll"
	"github.com/damoye/gopu/config"
	"github.com/damoye/gopu/httplog"
)

func main() {
	http.HandleFunc("/send", bll.HandleSend)
	http.HandleFunc("/receive", bll.HandleReceive)
	go bll.SendLoop()
	log.Fatal(http.ListenAndServe(
		config.Conf.HTTPAddress,
		httplog.New(http.DefaultServeMux),
	))
}
