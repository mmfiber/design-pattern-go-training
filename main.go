package main

type Executer interface {
	Key() string
	Title() string
	Do()
}

func main() {}
