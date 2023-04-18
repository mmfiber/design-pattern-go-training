package adapter

import "fmt"

type Printer interface {
	PrintWeak()
	PrintStrong()
}

type Banner struct {
	text string
}

func NewBanner() Banner {
	return Banner{}
}

func (b *Banner) ShowWithParen() {
	fmt.Printf("( %s )\n", b.text)
}

func (b *Banner) ShowWithAster() {
	fmt.Printf("* %s *\n", b.text)
}

type PrintBanner struct {
	banner *Banner
}

func NewPrintBanner(text string) PrintBanner {
	banner := Banner{text}
	return PrintBanner{&banner}
}

func (pb PrintBanner) PrintWeak() {
	pb.banner.ShowWithParen()
}

func (pb PrintBanner) PrintStrong() {
	pb.banner.ShowWithAster()
}

func main() {
	var p Printer
	p = NewPrintBanner("Hello")

	p.PrintWeak()
	p.PrintStrong()
}

// codes for main package
type Executer struct{}

func NewExecuter() Executer {
	return Executer{}
}

func (e Executer) Label() string {
	return "adapter pattern"
}

func (e Executer) Do() {
	main()
}
