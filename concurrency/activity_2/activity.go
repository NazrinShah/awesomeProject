package main

import (
	"errors"
	"fmt"
	"sync"
)

type BankBalance struct {
	amount   float64
	currency string
	mtx      sync.RWMutex
}

const (
	DISPLAY = iota + 1
	DEPOSIT
	WITHDRAW
	QUIT
)

var (
	Balance = &BankBalance{
		amount:   0,
		currency: "SGD",
	}

	InputLock = &sync.Mutex{}
)

func (bb *BankBalance) Deposit(amount float64) error {

	if amount < 0.0 {
		return errors.New("amount must be more than 0")
	}

	bb.mtx.Lock()
	bb.amount += amount
	bb.mtx.Unlock()

	return nil
}

func (bb *BankBalance) Withdraw(amount float64) error {

	bb.mtx.Lock()
	defer bb.mtx.Unlock()

	if amount > bb.amount {
		return errors.New("insufficient funds")
	}

	bb.amount -= amount

	return nil
}

func (bb *BankBalance) Display() {
	bb.mtx.RLock()
	defer bb.mtx.RUnlock()

	fmt.Println(fmt.Sprintf("Account has %.2f %s", bb.amount, bb.currency))
}

func processBankBalance(wg *sync.WaitGroup) {
	defer wg.Done()
	in := 0
	locked := false

	for in != QUIT {
		if !locked {
			locked = InputLock.TryLock()

			if !locked {
				continue
			}
		}

		fmt.Println("1. Display Account Balance\n2. Deposit into account\n3. Withdraw from account\n4. Quit")
		_, err := fmt.Scan(&in)

		if err != nil {
			fmt.Println("Error processing input: ", err)
			continue
		} else if in < DISPLAY || in > QUIT {
			fmt.Println(fmt.Sprintf("Invalid input. Expecting values between %d-%d but received %d", DISPLAY, QUIT, in))
			continue
		}

		switch in {
		case DISPLAY:
			Balance.Display()
		case DEPOSIT:
			for {
				amt := 0.0
				fmt.Printf("Enter amount to deposit: ")
				_, err := fmt.Scan(&amt)
				fmt.Println()

				if err != nil {
					fmt.Println("Error processing input: ", err)
					continue
				}

				err = Balance.Deposit(amt)

				if err != nil {
					fmt.Println("Error with deposit: ", err)
					continue
				}

				break
			}
		case WITHDRAW:
			amt := 0.0
			for {
				fmt.Printf("Enter amount to withdraw: ")
				_, err := fmt.Scan(&amt)
				fmt.Println()

				if err != nil {
					fmt.Println("Error processing input: ", err)
					continue
				}

				err = Balance.Withdraw(amt)
				if err != nil {
					fmt.Println("Error with withdrawal: ", err)
					continue
				}

				break
			}
		}

		InputLock.Unlock()
		locked = false
	}
}

func main() {
	wg := &sync.WaitGroup{}
	n := 2
	for i := 0; i < n; i++ {
		wg.Add(1)
		go processBankBalance(wg)
	}

	wg.Wait()
}
