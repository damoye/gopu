# gopu
Golang push server

### Server HTTP

url: `/send`

method: `POST`

request body:

    {
        "data": "data",
        "tokens": ["token1", "token2"]
    }

### Client WebSocket

url: `/receive?token=:token`

received message format: `{"task_id":100,"data":"data"}`

client should send task_id to server to confirm: `100`
