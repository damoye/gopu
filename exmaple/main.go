package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/receive?token=mytoken", nil)
	if err != nil {
		log.Fatal("dail:", err)
	}
	go func() {
		defer c.Close()
		for {
			_, p, err := c.ReadMessage()
			if err != nil {
				log.Fatal("readMessage:", err)
			}
			log.Printf("recv: %s", p)
			var message struct {
				TaskID int64  `json:"task_id"`
				Data   string `json:"data"`
			}
			if json.Unmarshal(p, &message) != nil {
				log.Fatal("unmarshal:", err)
			}
			err = c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprint(message.TaskID)))
			if err != nil {
				log.Fatal("writeMessage:", err)
			}
		}
	}()

	for {
		resp, err := http.Post(
			"http://localhost:8080/send", "application/json",
			bytes.NewBuffer([]byte("{\"tokens\":[\"mytoken\"],\"data\":\"data\"}")),
		)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp.Status)
		time.Sleep(time.Second)
	}
}
