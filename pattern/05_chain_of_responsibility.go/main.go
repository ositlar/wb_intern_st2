package main

import "fmt"

// Handler provides a handler interface.
type Handler interface {
	SendRequest(message int) int
}

// ConcreteHandlerA implements handler "A".
type ConcreteHandlerA struct {
	next Handler
}

// SendRequest implementation.
func (h *ConcreteHandlerA) SendRequest(message int) int {
	result := 0
	if message > 0 && message < 5 {
		fmt.Print("Handler 1: ")
		result = message * message
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return result
}

// ConcreteHandlerB implements handler "B".
type ConcreteHandlerB struct {
	next Handler
}

// SendRequest implementation.
func (h *ConcreteHandlerB) SendRequest(message int) int {
	result := 0
	if message >= 5 && message < 10 {
		result = message * message
		fmt.Print("Handler 2: ")
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return result
}

// ConcreteHandlerC implements handler "C".
type ConcreteHandlerC struct {
	next Handler
}

// SendRequest implementation.
func (h *ConcreteHandlerC) SendRequest(message int) int {
	result := 0
	if h.next == nil {
		result = message * message
		fmt.Print("Handler 3: ")
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return result
}

func main() {
	handlers := &ConcreteHandlerA{
		next: &ConcreteHandlerB{
			next: &ConcreteHandlerC{},
		},
	}
	for i := 1; i < 20; i++ {
		fmt.Println(handlers.SendRequest(i))
	}
}
