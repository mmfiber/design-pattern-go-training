package singleton

import (
	"fmt"
	"sync"
)

type User struct {
	id   string
	name string
}

type DBClient interface {
	GetUser() User
}

type MySqlDBClient struct{}

func (c *MySqlDBClient) GetUser() User {
	return User{"1", "sample user"}
}

var (
	dbClient DBClient
	lock     = &sync.Mutex{}
)

func getDBClient(ch chan DBClient) {
	if dbClient == nil {
		lock.Lock()
		defer lock.Unlock()

		fmt.Println("Cleate new instance")
		dbClient = &MySqlDBClient{}
	} else {
		fmt.Println("Instance is already decleared")
	}

	ch <- dbClient
}

func main() {
	ch := make(chan DBClient, 10)
	for i := 0; i < 30; i++ {
		go getDBClient(ch)
	}

	for cilent := range ch {
		fmt.Printf("client address: %p\n", &cilent)
	}
}

// codes for main package
type Executer struct{}

func NewExecuter() Executer {
	return Executer{}
}

func (e Executer) Label() string {
	return "singleton pattern"
}

func (e Executer) Do() {
	main()
}
