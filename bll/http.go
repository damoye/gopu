package bll

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/damoye/gopu/dal"
	"github.com/gorilla/websocket"
)

type sendRequest struct {
	Tokens []string `json:"tokens"`
	Data   string   `json:"data"`
}

func send(req *sendRequest) error {
	if len(req.Tokens) == 0 {
		return nil
	}
	taskID, err := dal.InsertTask(req.Data)
	if err != nil {
		return err
	}
	if err = dal.InsertSubtask(taskID, req.Tokens); err != nil {
		return err
	}
	b, err := json.Marshal(channelMessage{TaskID: taskID, Tokens: req.Tokens, Data: req.Data})
	if err != nil {
		panic(err)
	}
	return redisClient.Publish(redisChannel, string(b)).Err()
}

// HandleSend ...
func HandleSend(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var body sendRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := send(&body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// SendLoop ...
func SendLoop() {
	pubSub, err := redisClient.Subscribe(redisChannel)
	if err != nil {
		panic(err)
	}
	for {
		msg, err := pubSub.ReceiveMessage()
		var message channelMessage
		if err = json.Unmarshal([]byte(msg.Payload), &message); err != nil {
			log.Printf("unmarshal redisMessage %q: %v", msg.Payload, err)
			continue
		}
		bytes, err := json.Marshal(struct {
			TaskID int64  `json:"task_id"`
			Data   string `json:"data"`
		}{TaskID: message.TaskID, Data: message.Data})
		if err != nil {
			panic(err)
		}
		for _, conn := range clients.MGet(message.Tokens) {
			if err = conn.WriteMessage(websocket.TextMessage, bytes); err == nil {
				continue
			}
			log.Println("write websocket:", err)
			if err = conn.Close(); err != nil {
				log.Println("close websocket:", err)
			}
		}
	}
}
