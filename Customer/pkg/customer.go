package pkg

import "fmt"

type Customer struct {
	fName    string
	lName    string
	username string
	password string
	email    string
	phone    int
	address  string
}

func CreateCustomer(fName, lName, username, password, email string, phone int, address string) *Customer {
	return &Customer{
		fName:    fName,
		lName:    lName,
		username: username,
		password: password,
		email:    email,
		phone:    phone,
		address:  address,
	}
}

func (c *Customer) RetrieveUserCredentials() (string, string) {
	return c.username, c.password
}

func (c *Customer) RetrieveUserAddress() string {
	return c.address
}

func (c *Customer) PrintUserInformation() {
	fmt.Println(fmt.Sprintf("First Name: %s\nLast Name: %s\nUsername: %s\nPassword: %sEmail: %s\nPhone: %d\nAddress: %s",
		c.fName, c.lName, c.username, c.password, c.email, c.phone, c.address))
}
