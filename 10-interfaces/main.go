package main

import (
	"fmt"
)

type channelchanger interface {
	upChannel()
	downChannel()
}

// TV aka Television
type TV struct {
	channel int
}

// Remote TV remote control
type Remote struct {
	tv *TV
}

func (r *Remote) sync(tv *TV) {
	r.tv = tv
}

func (r *Remote) upChannel() {
	r.tv.channel++
	fmt.Printf("Remote up %v\n", r.tv.channel)
}

func (r *Remote) downChannel() {
	r.tv.channel--
	fmt.Printf("Remote down %v\n", r.tv.channel)
}

// Buttonpad TV buttons
type Buttonpad struct {
	tv *TV
}

func (b *Buttonpad) attach(tv *TV) {
	b.tv = tv
}

func (b *Buttonpad) upChannel() {
	b.tv.channel++
	fmt.Printf("Buttonpad up %v\n", b.tv.channel)
}

func (b *Buttonpad) downChannel() {
	b.tv.channel--
	fmt.Printf("Buttonpad down %v\n", b.tv.channel)
}

type control interface {
	upChannel()
	downChannel()
}

func upChannel(c control) {
	c.upChannel()
}

func downChannel(c control) {
	c.downChannel()
}

func main() {
	tv := TV{}
	buttonpad := Buttonpad{}
	buttonpad.attach(&tv)

	remote := Remote{}
	remote.sync(&tv)

	fmt.Println("TV on")
	fmt.Printf("Default channel is %v\n", tv.channel)

	upChannel(&remote)
	upChannel(&buttonpad)
	upChannel(&buttonpad)
	upChannel(&remote)
	upChannel(&remote)
	upChannel(&buttonpad)
	downChannel(&remote)
	downChannel(&remote)
	downChannel(&remote)
	downChannel(&remote)
	downChannel(&buttonpad)
}
