package flyweight

import "fmt"

type FlyweightFactory struct {
	flyweights map[string]Flyweight
}

type Flyweight interface {
	Operation(extrinsicState string)
}

type ConcreteFlyweight struct {
	intrinsicState string
}

func (cf *ConcreteFlyweight) Operation(extrinsicState string) {
	fmt.Printf("Flyweight Address: %p, Intrinsic State: %s, Extrinsic State: %s\n", cf, cf.intrinsicState, extrinsicState)
}

func (ff *FlyweightFactory) GetFlyweight(key string) Flyweight {
	// すでに存在するFlyweightオブジェクトを再利用
	if flyweight, ok := ff.flyweights[key]; ok {
		return flyweight
	}

	flyweight := &ConcreteFlyweight{intrinsicState: key}
	ff.flyweights[key] = flyweight
	return flyweight
}

func main() {
	factory := FlyweightFactory{
		flyweights: make(map[string]Flyweight),
	}

	// Flyweightオブジェクトの取得と操作
	flyweight1 := factory.GetFlyweight("A")
	flyweight1.Operation("state1")

	flyweight2 := factory.GetFlyweight("B")
	flyweight2.Operation("state2")

	flyweight3 := factory.GetFlyweight("A")
	flyweight3.Operation("state3")
}

// codes for main package
type Executer struct{}

func NewExecuter() Executer {
	return Executer{}
}

func (e Executer) Label() string {
	return "flyweight pattern"
}

func (e Executer) Do() {
	main()
}
