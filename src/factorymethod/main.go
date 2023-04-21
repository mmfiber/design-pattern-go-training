package factorymethod

import "fmt"

type Product interface {
	Use()
}

type ProductCreator[T Product] interface {
	createProduct(owner Owner) T
	registerProduct(product T)
}

type Factory[T Product] struct {
	ProductCreator[T]
}

func NewFactory[T Product](pc ProductCreator[T]) Factory[T] {
	return Factory[T]{pc}
}

func (f Factory[T]) Create(owner Owner) T {
	product := f.createProduct(owner)
	f.registerProduct(product)
	return product
}

type Owner struct {
	name string
}

func (o *Owner) Name() string {
	return o.name
}

type IDCard struct {
	owner *Owner
}

func NewIDCard(owner Owner) IDCard {
	fmt.Printf("create card (owner: %s)\n", owner.Name())
	return IDCard{&owner}
}

func (c *IDCard) Owner() Owner {
	return *c.owner
}

func (c *IDCard) Use() {
	fmt.Printf("use card (owner: %s)\n", c.owner.Name())
}

type IDCardCreator struct {
	owners []*Owner
}

func NewIDCardCreator() IDCardCreator {
	return IDCardCreator{[]*Owner{}}
}

func (f *IDCardCreator) createProduct(owner Owner) *IDCard {
	idCard := NewIDCard(owner)
	return &idCard
}

func (f *IDCardCreator) registerProduct(product *IDCard) {
	f.owners = append(f.owners, product.owner)
}

func main() {
	idCardCreator := NewIDCardCreator()
	idCardFactory := NewFactory[*IDCard](&idCardCreator)

	owners := []Owner{
		Owner{"デク"},
		Owner{"かっちゃん"},
		Owner{"ショート"},
	}
	for _, owner := range owners {
		card := idCardFactory.Create(owner)
		card.Use()
	}
}

// codes for main package
type Executer struct{}

func NewExecuter() Executer {
	return Executer{}
}

func (e Executer) Label() string {
	return "factory method pattern"
}

func (e Executer) Do() {
	main()
}
