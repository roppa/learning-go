package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// Message a socket message
type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

// Channel is a socket channel
type Channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4000", nil)
}

func addChannel(data interface{}) {
	var channel Channel
	channelMap := data.(map[string]interface{})
	channel.Name = channelMap["name"].(string)
	channel.ID = channelMap["id"].(string)
	fmt.Printf("%v#\n", channel)
}

func subscribeChannel(socket *websocket.Conn) {
	for {
		time.Sleep(time.Second * 1)
		message := Message{
			"channel add",
			Channel{"1", "software support"},
		}
		socket.WriteJSON(message)
		fmt.Println("sent new channel")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		var inMessage Message
		if err := socket.ReadJSON(&inMessage); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%#v\n", inMessage)

		switch inMessage.Name {
		case "channel add":
			addChannel(inMessage.Data)
		case "channel subscribe":
			go subscribeChannel(socket)
		}
	}
}
