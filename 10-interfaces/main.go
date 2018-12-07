package main

import (
	"fmt"
)

// TV aka Television
type TV struct {
	channel int
}

func (tv *TV) upChannel() {
	tv.channel++
}

func (tv *TV) downChannel() {
	tv.channel--
}

// Remote TV remote control
type Remote struct {
	tv *TV
}

func (r *Remote) up(tv *TV) {
	tv.upChannel()
}

// Buttonpad TV buttons
type Buttonpad struct {
	tv *TV
}

type control interface {
	upChannel()
	downChannel()
}

func main() {
	tv := TV{}
	fmt.Println("TV on")
	fmt.Printf("Default channel is %v\n", tv.channel)
	tv.upChannel()
	fmt.Printf("Channel is %v\n", tv.channel)
}
