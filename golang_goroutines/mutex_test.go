package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutexcondition(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}

type Bankaccount struct {
	RWmutex sync.RWMutex
	balance int
}

func (account *Bankaccount) Addbalance(amount int) {
	account.RWmutex.Lock()
	account.balance = account.balance + amount
	account.RWmutex.Unlock()
}

func (account *Bankaccount) Getbalance() int {
	account.RWmutex.RLock()
	balance := account.balance
	account.RWmutex.RUnlock()
	return balance
}

func TestRWmutex(t *testing.T) {
	account := Bankaccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.Addbalance(1)
				fmt.Println(account.Getbalance())
			}
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("total balance = ", account.Getbalance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("lock user1", user1.Name)
	user1.Change(-amount)
	time.Sleep(2 * time.Second)

	user2.Lock()
	fmt.Println("lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(2 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "sherdhan",
		Balance: 1000,
	}
	user2 := UserBalance{
		Name:    "syarif",
		Balance: 1000,
	}

	go Transfer(&user1, &user2, 1000)

	time.Sleep(2 * time.Second)

	fmt.Println("user = ", user1.Name, ", balance = ", user1.Balance)
	fmt.Println("user = ", user2.Name, ", balance = ", user2.Balance)
}
