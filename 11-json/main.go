package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	recRawMessage := []byte(`{"name": "channel add", "data": { "name": "First channel" }}`)
	var recMessage Message
	err := json.Unmarshal(recRawMessage, &recMessage)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", recMessage)

	if recMessage.Name == "channel add" {
		channel, err := addChannel(recMessage.Data)
		var sendMessage Message
		sendMessage.Name = "channel add"
		sendMessage.Data = channel
		sendRawMsg, err := json.Marshal(sendMessage)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf(string(sendRawMsg))
	}
}

func addChannel(data interface{}) (Channel, error) {
	var channel Channel
	channelMap := data.(map[string]interface{})
	channel.Name = channelMap["name"].(string)
	channel.ID = "1"
	fmt.Printf("%v#\n", channel)
	return channel, nil
}
