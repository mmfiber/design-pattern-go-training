package decorator

import "fmt"

type Display interface {
	GetColumns() int
	GetRows() int
	GetRowText(row int) (string, bool)
}

type StringDisplay struct {
	value string
}

func NewStringDisplay(value string) *StringDisplay {
	return &StringDisplay{value}
}

func (d *StringDisplay) GetColumns() int {
	return len(d.value)
}

func (d *StringDisplay) GetRows() int {
	return 1
}

func (d *StringDisplay) GetRowText(row int) (string, bool) {
	if row != 0 {
		return "", false
	}
	return d.value, true
}

func main() {
	show := func(display Display) {
		for i := 0; i < display.GetRows(); i++ {
			if text, ok := display.GetRowText(i); ok {
				fmt.Println(text)
			}
		}
	}

	b1 := NewStringDisplay("Hello, world.")
	b2 := NewSideBorder(b1, "#")
	b3 := NewFullBorder(b2)
	show(b1)
	show(b2)
	show(b3)

	b4 := NewSideBorder(
		NewFullBorder(
			NewFullBorder(
				NewSideBorder(
					NewFullBorder(b1),
					"*",
				),
			),
		),
		"/",
	)
	show(b4)
}

// codes for main package
type Executer struct{}

func NewExecuter() Executer {
	return Executer{}
}

func (e Executer) Label() string {
	return "decorator pattern"
}

func (e Executer) Do() {
	main()
}
