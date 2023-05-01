package strategy

import (
	"fmt"
)

type RouteStrategy interface {
	Route(from, to string)
}

type WalkingStrategy struct{}

func (s *WalkingStrategy) Route(from, to string) {
	fmt.Printf("Walking route from %s to %s: 4 km, 30 min\n", from, to)
}

type PublicTransportStrategy struct{}

func (s *PublicTransportStrategy) Route(from, to string) {
	fmt.Printf("Public transport route from %s to %s: 3 km, 5 min\n", from, to)
}

type Navigator struct {
	strategy RouteStrategy
}

func (n *Navigator) Route(from, to string) {
	n.strategy.Route(from, to)
}

func main() {
	var navigator *Navigator

	navigator = &Navigator{&WalkingStrategy{}}
	navigator.Route("Shibuya", "Yoyogi")
	navigator.Route("Yoyogi", "Shibuya")

	navigator = &Navigator{&PublicTransportStrategy{}}
	navigator.Route("Shibuya", "Yoyogi")
	navigator.Route("Yoyogi", "Shibuya")
}

// codes for main package
type Executer struct{}

func NewExecuter() Executer {
	return Executer{}
}

func (e Executer) Label() string {
	return "starategy pattern"
}

func (e Executer) Do() {
	main()
}
