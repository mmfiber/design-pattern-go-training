package facade

import "fmt"

type OrderService struct{}

func (os *OrderService) PlaceOrder() {
	fmt.Println("注文を受け付けました")
}

type PaymentService struct{}

func (ps *PaymentService) ProcessPayment() {
	fmt.Println("支払いを処理しました")
}

type ShippingService struct{}

func (ss *ShippingService) ShipProduct() {
	fmt.Println("商品を発送しました")
}

type OrderFacade struct {
	orderService    *OrderService
	paymentService  *PaymentService
	shippingService *ShippingService
}

func (of *OrderFacade) PlaceOrder() {
	of.orderService.PlaceOrder()
	of.paymentService.ProcessPayment()
	of.shippingService.ShipProduct()
}

func main() {
	orderFacade := &OrderFacade{
		orderService:    &OrderService{},
		paymentService:  &PaymentService{},
		shippingService: &ShippingService{},
	}

	orderFacade.PlaceOrder()
}

// codes for main package
type Executer struct{}

func NewExecuter() Executer {
	return Executer{}
}

func (e Executer) Label() string {
	return "facade pattern"
}

func (e Executer) Do() {
	main()
}
