package main

import "sync"

type BankAccount struct {
	balance int
	mutex   sync.Mutex
}

func (ba *BankAccount) Deposit(amount int, wg *sync.WaitGroup) {
	ba.mutex.Lock()
	ba.balance = ba.balance + amount
	ba.mutex.Unlock()
	wg.Done()
}

func (ba *BankAccount) Withdraw(amount int, wg *sync.WaitGroup) {
	ba.mutex.Lock()
	ba.balance = ba.balance - amount
	ba.mutex.Unlock()
	wg.Done()
}

func bankAccount() {
	acc1 := BankAccount{balance: 1000}
	var wg sync.WaitGroup
	wg.Add(6)
	go acc1.Deposit(100, &wg)
	go acc1.Deposit(100, &wg)
	go acc1.Withdraw(100, &wg)
	go acc1.Deposit(100, &wg)
	go acc1.Withdraw(100, &wg)
	go acc1.Deposit(100, &wg)

	wg.Wait()
	println(acc1.balance)
}
