package abstractfactory

import "fmt"

type MacBtn struct{}

func (b *MacBtn) Click() {
	fmt.Println("Mac btn")
}

type MacUIFactory struct{}

func NewMacUIFactory() UIFactory {
	return &MacUIFactory{}
}

func (f *MacUIFactory) CreateBtn() Btn {
	return &MacBtn{}
}
