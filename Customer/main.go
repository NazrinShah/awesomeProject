package main

import (
	"fmt"
	"github.com/NazrinShah/awesomeProject/Customer/pkg"
)

func main() {
	Customer1 := pkg.CreateCustomer("Michael", "Jordan", "MJ2020", "1234567", "MJ2020@gmail.com", 12345678, "18227 Capstan Greens Road Cornelius, NC 28031")
	Customer1.PrintUserInformation()
	fmt.Println()

	u, p := Customer1.RetrieveUserCredentials()
	a := Customer1.RetrieveUserAddress()
	fmt.Println(fmt.Sprintf("Username: %s\nPassword: %s\n", u, p))
	fmt.Println("Address: ", a)
}
