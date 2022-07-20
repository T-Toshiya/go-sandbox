package main

import (
	"fmt"
)

type Message string

type Greeter struct {
	Message Message
}

type Event struct {
	Greeter Greeter
}

func NewMessage() Message {
	return "Hi there!"
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (g Greeter) Greet() Message {
	return g.Message
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

//func main() {
//	message := NewMessage()
//	greeter := NewGreeter(message)
//	event := NewEvent(greeter)
//
//	event.Start()
//}

func main() {
	e := InitializeEvent()

	e.Start()
}
