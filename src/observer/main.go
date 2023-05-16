package observer

import "fmt"

type Subject interface {
	RegisterObserver(Observer)
	UnregisterObserver(Observer)
	NotifyObservers()
}

type Observer interface {
	Update(Subject)
}

type ConcreteSubject struct {
	observers []Observer
}

func (s *ConcreteSubject) RegisterObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) UnregisterObserver(observer Observer) {
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) NotifyObservers() {
	for _, observer := range s.observers {
		observer.Update(s)
	}
}

type ConcreteObserver struct {
	name string
}

func (o *ConcreteObserver) Update(subject Subject) {
	fmt.Printf("Observer %s received notification from Subject\n", o.name)
}

func main() {
	obs1 := &ConcreteObserver{name: "obserber1"}
	obs2 := &ConcreteObserver{name: "obserber2"}

	sub := &ConcreteSubject{}
	sub.RegisterObserver(obs1)
	sub.RegisterObserver(obs2)

	sub.NotifyObservers()

	sub.UnregisterObserver(obs1)

	sub.NotifyObservers()
}

// codes for main package
type Executer struct{}

func NewExecuter() Executer {
	return Executer{}
}

func (e Executer) Label() string {
	return "obserber pattern"
}

func (e Executer) Do() {
	main()
}
