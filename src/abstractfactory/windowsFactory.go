package abstractfactory

import "fmt"

type WindowsBtn struct{}

func (b *WindowsBtn) Click() {
	fmt.Println("windows btn")
}

type WindowsUIFactory struct{}

func NewWindowsUIFactory() UIFactory {
	return &WindowsUIFactory{}
}

func (f *WindowsUIFactory) CreateBtn() Btn {
	return &WindowsBtn{}
}
