package state

import "fmt"

type State interface {
	InsertCoin()
	SelectProduct()
	Dispense()
}

type NoCoinState struct{}

func (s *NoCoinState) InsertCoin() {
	fmt.Println("コインが挿入されました。")
}

func (s *NoCoinState) SelectProduct() {
	// error
}

func (s *NoCoinState) Dispense() {
	// error
}

type HasCoinState struct{}

func (s *HasCoinState) InsertCoin() {
	fmt.Println("既にコインが挿入されています。")
}

func (s *HasCoinState) SelectProduct() {
	fmt.Println("商品が選択されました。")
}

func (s *HasCoinState) Dispense() {
	fmt.Println("商品が提供されました。")
}

type VendingMachine struct {
	state State
}

func (vm *VendingMachine) InsertCoin() {
	vm.state.InsertCoin()
	vm.state = &HasCoinState{}
}

func (vm *VendingMachine) SelectProduct() {
	vm.state.SelectProduct()
}

func (vm *VendingMachine) Dispense() {
	vm.state.Dispense()
	vm.state = &NoCoinState{}
}

func main() {
	vendingMachine := &VendingMachine{state: &NoCoinState{}}

	vendingMachine.InsertCoin()    // コインが挿入されました。
	vendingMachine.SelectProduct() // 商品が選択されました。
	vendingMachine.Dispense()      // 商品が提供されました。
}

// codes for main package
type Executer struct{}

func NewExecuter() Executer {
	return Executer{}
}

func (e Executer) Label() string {
	return "state pattern"
}

func (e Executer) Do() {
	main()
}
