package main

import "errors"

// ------------------------------ Initial Constructor -------------------------------------------------- //
func newMessage(msg string) message {
	return message{value: msg}
}

func newGreeter(m message) greeter {
	return greeter{message: m}
}

func newEvent(g greeter) (event, error) {
	//return event{greeter: g}
	return event{greeter: g}, errors.New("Failed event")
}
