package visitor

import (
	"fmt"
	"math"
)

type Visitor interface {
	VisitForSquare(*Square)
	VisitForCircle(*Circle)
	VisitForRectangle(*Rectangle)
}

type AreaCalculator struct {
	area float64
}

func (a *AreaCalculator) VisitForSquare(s *Square) {
	fmt.Println("Calculating area for square")
	a.area = s.side * s.side
}

func (a *AreaCalculator) VisitForCircle(s *Circle) {
	fmt.Println("Calculating area for circle")
	a.area = s.radius * s.radius * math.Pi
}

func (a *AreaCalculator) VisitForRectangle(s *Rectangle) {
	fmt.Println("Calculating area for rectangle")
	a.area = s.l * s.b
}

func (a *AreaCalculator) CalculatedArea() {
	fmt.Printf("Calculated area is: %.2f\n", a.area)
}

func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}

	square.Accept(areaCalculator)
	areaCalculator.CalculatedArea()

	circle.Accept(areaCalculator)
	areaCalculator.CalculatedArea()

	rectangle.Accept(areaCalculator)
	areaCalculator.CalculatedArea()
}

// codes for main package
type Executer struct{}

func NewExecuter() Executer {
	return Executer{}
}

func (e Executer) Label() string {
	return "visitor pattern"
}

func (e Executer) Do() {
	main()
}
