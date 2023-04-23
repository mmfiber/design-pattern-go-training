package main

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/mmfiber/design-pattern-go-training/src/adapter"
	"github.com/mmfiber/design-pattern-go-training/src/exit"
	"github.com/mmfiber/design-pattern-go-training/src/factorymethod"
	"github.com/mmfiber/design-pattern-go-training/src/iterator"
	"github.com/mmfiber/design-pattern-go-training/src/singleton"
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
	}

	for {
		idx, _, err := prompt.Run()
		if err != nil {
			log.Fatal(fmt.Printf("Prompt failed %v\n", err))
		}
		executers[idx].Do()
	}
}
