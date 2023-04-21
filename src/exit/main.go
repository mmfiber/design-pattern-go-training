package exit

import "os"

type Executer struct{}

func NewExecuter() Executer {
	return Executer{}
}

func (e Executer) Label() string {
	return "exit"
}

func (e Executer) Do() {
	os.Exit(0)
}
