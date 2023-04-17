package main

import (
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/mmfiber/design-pattern-go-training/src/iterator"
)

type Executer interface {
	Label() string
	Do()
}

type ExecuterItems struct {
	idx   int
	Label string
}

func main() {
	items := []ExecuterItems{}
	executers := []Executer{iterator.NewExecuter()}
	for idx, executer := range executers {
		items = append(items, ExecuterItems{idx, executer.Label()})
	}

	exitItem := ExecuterItems{len(items), "exit"}
	prompt := promptui.Select{
		Label: "Select",
		Items: append(items, exitItem),
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
		if idx == exitItem.idx {
			os.Exit(0)
		}
		executers[idx].Do()
	}
}
