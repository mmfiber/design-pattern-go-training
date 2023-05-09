package chainofresponsibility

import (
	"fmt"
)

type Handler interface {
	next() Handler
	setNext(Handler)
	handle()
}

type ConcreteHandler struct {
	id      string
	handler Handler
}

func NewConcreteHandler(id string) *ConcreteHandler {
	return &ConcreteHandler{id, nil}
}

func (h *ConcreteHandler) next() Handler {
	return h.handler
}

func (h *ConcreteHandler) setNext(handler Handler) {
	h.handler = handler
}

func (h *ConcreteHandler) handle() {
	fmt.Printf("Handler %s\n", h.id)
	next := h.next()
	if next == nil {
		return
	}
	next.handle()
}

func main() {
	h1 := NewConcreteHandler("1")
	h2 := NewConcreteHandler("2")
	h3 := NewConcreteHandler("3")

	fmt.Println("Process asc order...")
	h1.setNext(h2)
	h2.setNext(h3)
	h1.handle()

	fmt.Println("Process desc order...")
	h3.setNext(h2)
	h2.setNext(h1)
	h1.setNext(nil)
	h3.handle()

	fmt.Println("Process handler which does not have next...")
	h1.handle()
}

// codes for main package
type Executer struct{}

func NewExecuter() Executer {
	return Executer{}
}

func (e Executer) Label() string {
	return "chain of responsibility pattern"
}

func (e Executer) Do() {
	main()
}
