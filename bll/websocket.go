package bll

import (
	"log"
	"net/http"
	"strconv"

	"github.com/damoye/gopu/dal"
	"github.com/damoye/gopu/utils"
	"github.com/gorilla/websocket"
)

var (
	clients  = utils.NewWebsocketConnMap()
	upgrader = websocket.Upgrader{}
)

// HandleReceive ...
func HandleReceive(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("websocket upgrade:", err)
		return
	}
	defer c.Close()
	token := r.URL.Query().Get("token")
	if token == "" {
		log.Println("no token")
		return
	}
	if !clients.SetNX(token, c) {
		log.Println("token exist:", token)
		return
	}
	defer clients.Del(token)
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("websocket readMessage:", err)
			break
		}
		taskID, err := strconv.ParseInt(string(message), 10, 64)
		if err != nil {
			log.Println("message parse:", err)
			break
		}
		if err = dal.Deliver(taskID, token); err != nil {
			log.Println("dal deliver:", err)
		}
	}
}
