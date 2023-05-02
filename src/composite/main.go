package composite

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

type Component interface {
	Operation() string
}

type Leaf struct {
	name string
}

func (l *Leaf) Operation() string {
	return l.name
}

type Composite struct {
	children []Component
	name     string
}

func (c *Composite) Add(component Component) {
	c.children = append(c.children, component)
}

func (c *Composite) Remove(component Component) {
	idx := slices.Index(c.children, component)
	if idx == -1 {
		return
	}

	c.children = append(c.children[:idx], c.children[idx+1:]...)
}

func (c *Composite) Operation() string {
	var results []string
	for _, c := range c.children {
		results = append(results, c.Operation())
	}
	return fmt.Sprintf("%s(%s)", c.name, strings.Join(results, " + "))
}

func main() {
	tree := &Composite{name: "Tree"}

	branch1 := &Composite{name: "Branch1"}
	branch1.Add(&Leaf{"Leaf1"})
	branch1.Add(&Leaf{"Leaf2"})

	branch2 := &Composite{name: "Branch2"}
	branch2.Add(&Leaf{"Leaf3"})

	tree.Add(branch1)
	tree.Add(branch2)

	fmt.Println(tree.Operation())

	fmt.Printf("\n// Remove Branch 1 from Tree...\n\n")
	tree.Remove(branch1)

	fmt.Println(tree.Operation())
}

// codes for main package
type Executer struct{}

func NewExecuter() Executer {
	return Executer{}
}

func (e Executer) Label() string {
	return "composite pattern"
}

func (e Executer) Do() {
	main()
}
