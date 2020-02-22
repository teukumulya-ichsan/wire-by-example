package main

import (
	"fmt"
	"log"
)

func main() {
	msg := "Welcome Gopher Ichsan"
	e, err := initializeEvent(msg)
	if err != nil {
		log.Fatalf("FAILURE: %v", err)
	}

	e.Start()
}

// ------------------------------ Object and Method -------------------------------------------------- //

type message struct {
	value string
}

func (m message) String() string {
	return fmt.Sprintf("[OFFICIAL MESSAGE]: %s\n", m.value)
}

type greeter struct {
	message message
}

func (g greeter) great() string {
	return g.message.String()
}

type event struct {
	greeter greeter
}

func (e event) Start() {
	fmt.Println(e.greeter.great())
}
