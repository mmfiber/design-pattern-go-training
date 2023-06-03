package main

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/mmfiber/design-pattern-go-training/src/abstractfactory"
	"github.com/mmfiber/design-pattern-go-training/src/adapter"
	"github.com/mmfiber/design-pattern-go-training/src/bridge"
	"github.com/mmfiber/design-pattern-go-training/src/builder"
	"github.com/mmfiber/design-pattern-go-training/src/chainofresponsibility"
	"github.com/mmfiber/design-pattern-go-training/src/composite"
	"github.com/mmfiber/design-pattern-go-training/src/decorator"
	"github.com/mmfiber/design-pattern-go-training/src/exit"
	"github.com/mmfiber/design-pattern-go-training/src/facade"
	"github.com/mmfiber/design-pattern-go-training/src/factorymethod"
	"github.com/mmfiber/design-pattern-go-training/src/flyweight"
	"github.com/mmfiber/design-pattern-go-training/src/iterator"
	"github.com/mmfiber/design-pattern-go-training/src/memento"
	"github.com/mmfiber/design-pattern-go-training/src/observer"
	"github.com/mmfiber/design-pattern-go-training/src/prototype"
	"github.com/mmfiber/design-pattern-go-training/src/proxy"
	"github.com/mmfiber/design-pattern-go-training/src/singleton"
	"github.com/mmfiber/design-pattern-go-training/src/state"
	"github.com/mmfiber/design-pattern-go-training/src/strategy"
	"github.com/mmfiber/design-pattern-go-training/src/visitor"
)

type Executer interface {
	Label() string
	Do()
}

func main() {
	executers := []Executer{
		iterator.NewExecuter(),
		adapter.NewExecuter(),
		factorymethod.NewExecuter(),
		singleton.NewExecuter(),
		prototype.NewExecuter(),
		builder.NewExecuter(),
		abstractfactory.NewExecuter(),
		bridge.NewExecuter(),
		strategy.NewExecuter(),
		composite.NewExecuter(),
		decorator.NewExecuter(),
		visitor.NewExecuter(),
		chainofresponsibility.NewExecuter(),
		facade.NewExecuter(),
		observer.NewExecuter(),
		memento.NewExecuter(),
		state.NewExecuter(),
		flyweight.NewExecuter(),
		proxy.NewExecuter(),
		exit.NewExecuter(),
	}
	prompt := promptui.Select{
		Label: "Select",
		Items: executers,
		Templates: &promptui.SelectTemplates{
			Active:   "\U0000276F {{ .Label | cyan }}",
			Inactive: "  {{ .Label | cyan }}",
			Selected: "{{ .Label | faint }}",
		},
		Size: len(executers),
	}

	for {
		idx, _, err := prompt.Run()
		if err != nil {
			log.Fatal(fmt.Printf("Prompt failed %v\n", err))
		}
		executers[idx].Do()
	}
}
