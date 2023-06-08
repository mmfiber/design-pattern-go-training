package interpreter

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type Expression interface {
	Interpret() bool
}

type TrueExpression struct{}

func (t *TrueExpression) Interpret() bool {
	return true
}

type FalseExpression struct{}

func (f *FalseExpression) Interpret() bool {
	return false
}

type AndExpression struct {
	left, right Expression
}

func (a *AndExpression) Interpret() bool {
	return a.left.Interpret() && a.right.Interpret()
}

type OrExpression struct {
	left, right Expression
}

func (o *OrExpression) Interpret() bool {
	return o.left.Interpret() || o.right.Interpret()
}

func Parse(tokens []string) Expression {
	if len(tokens) == 0 {
		return nil
	}

	switch tokens[0] {
	case "TRUE":
		return &TrueExpression{}
	case "FALSE":
		return &FalseExpression{}
	case "AND":
		return &AndExpression{Parse(tokens[1:]), Parse(tokens[2:])}
	case "OR":
		return &OrExpression{Parse(tokens[1:]), Parse(tokens[2:])}
	}

	return nil
}

func main() {
	// This represents the rule "TRUE AND (FALSE OR TRUE)"
	tokens := strings.Split("AND TRUE OR FALSE TRUE", " ")
	expression := Parse(tokens)
	spew.Printf("%#v\n", expression)
	// Output:
	// (*interpreter.AndExpression){left:(*interpreter.TrueExpression){} right:(*interpreter.OrExpression){
	//	left:(*interpreter.FalseExpression){} right:(*interpreter.TrueExpression){}
	// }}

	fmt.Println(expression.Interpret())
	// Output: true
}

// codes for main package
type Executer struct{}

func NewExecuter() Executer {
	return Executer{}
}

func (e Executer) Label() string {
	return "interpreter pattern"
}

func (e Executer) Do() {
	main()
}
