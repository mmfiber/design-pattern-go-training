package abstractfactory

type Btn interface {
	Click()
}

type UIFactory interface {
	CreateBtn() Btn
}

type OS int

const (
	Windwos OS = iota
	Mac
)

func GetUIFactory(os OS) UIFactory {
	switch os {
	case Windwos:
		return NewWindowsUIFactory()
	case Mac:
		return NewMacUIFactory()
	default:
		panic("invalid os")
	}
}

func main() {
	var factory UIFactory

	factory = GetUIFactory(Windwos)
	factory.CreateBtn().Click()

	factory = GetUIFactory(Mac)
	factory.CreateBtn().Click()
}

// codes for main package
type Executer struct{}

func NewExecuter() Executer {
	return Executer{}
}

func (e Executer) Label() string {
	return "abstract factory pattern"
}

func (e Executer) Do() {
	main()
}
