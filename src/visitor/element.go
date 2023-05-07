package visitor

type Shape interface {
	GetType() string
	Accept(Visitor)
}

type Square struct {
	side float64
}

func (s *Square) GetType() string {
	return "Square"
}

func (s *Square) Accept(v Visitor) {
	v.VisitForSquare(s)
}

type Circle struct {
	radius float64
}

func (c *Circle) GetType() string {
	return "Circle"
}

func (c *Circle) Accept(v Visitor) {
	v.VisitForCircle(c)
}

type Rectangle struct {
	l float64
	b float64
}

func (t *Rectangle) GetType() string {
	return "rectangle"
}

func (t *Rectangle) Accept(v Visitor) {
	v.VisitForRectangle(t)
}
