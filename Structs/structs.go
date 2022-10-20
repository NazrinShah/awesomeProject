package main

import "fmt"

type User struct {
	FName            string
	LName            string
	Age              int
	Subscriber       bool
	HomeAddress      string
	Phone            int
	CreditAvailable  float32
	CurrentCartCost  float32
	CurrentOrderCost float32
}

func (u *User) PrintDetails() {
	details := ""

	details = fmt.Sprintf("First Name: %s", u.FName)
	details = fmt.Sprintf("%s\nLast Name: %s", details, u.LName)
	details = fmt.Sprintf("%s\nAge: %d", details, u.Age)

	temp := "yes"

	if !u.Subscriber {
		temp = "no"
	}

	details = fmt.Sprintf("%s\nSubscriber: %s", details, temp)
	details = fmt.Sprintf("%s\nHome Address: %s", details, u.HomeAddress)
	details = fmt.Sprintf("%s\nPhone: %d", details, u.Phone)
	details = fmt.Sprintf("%s\nCredit Available: %.2f", details, u.CreditAvailable)
	details = fmt.Sprintf("%s\nCurrent Cart Cost: %.2f", details, u.CurrentCartCost)
	details = fmt.Sprintf("%s\nCurrent Order Cost: %.2f", details, u.CurrentOrderCost)

	fmt.Println(details)
}

func main() {
	c1 := User{
		FName:            "Annakin",
		LName:            "Skywalker",
		Age:              45,
		Subscriber:       true,
		HomeAddress:      "Death Star",
		Phone:            1234567,
		CreditAvailable:  10000.00,
		CurrentCartCost:  0,
		CurrentOrderCost: 0,
	}

	c2 := User{
		FName:            "Han",
		LName:            "Solo",
		Age:              50,
		Subscriber:       false,
		HomeAddress:      "Tatooine",
		Phone:            4321765,
		CreditAvailable:  20000,
		CurrentCartCost:  0,
		CurrentOrderCost: 0,
	}

	c1.PrintDetails()
	fmt.Println()
	c2.PrintDetails()
}
