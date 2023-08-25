package main

import (
	"fmt"
	"sync"
)

var (
	mu      sync.Mutex
	balance int
)

func deposit(val int, wg *sync.WaitGroup) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Printf("🟡 Depositing $%d to account with current balance $%d\n", val, balance)
	balance += val
	fmt.Printf("✅ Deposit success! Updated balance: $%d\n\n", balance)

	wg.Done()
}

func withdraw(val int, wg *sync.WaitGroup) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Printf("🔴 Withdrawing $%d from account with current balance $%d\n", val, balance)
	balance -= val
	fmt.Printf("✅ Withdrawal success! Updated balance: $%d\n\n", balance)

	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	balance = 1000

	wg.Add(2)
	go withdraw(700, &wg)
	go deposit(500, &wg)

	wg.Wait()

	fmt.Printf("Final balance: $%d", balance)
}
