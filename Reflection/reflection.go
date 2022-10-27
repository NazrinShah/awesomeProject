package main

import (
	"fmt"
	"reflect"
)

type Customer struct {
	FName        string
	LName        string
	UserID       int
	InvoiceTotal float64
}

func inspect(in interface{}) {
	switch in.(type) {
	case Customer:
		v := reflect.ValueOf(in)
		fmt.Println(fmt.Sprintf("Customer has %d fields", v.NumField()))

		for i := 0; i < v.NumField(); i++ {
			fmt.Println(fmt.Sprintf("type %s has value [%+v]", v.Field(i).Type().Name(), v.Field(i).Interface()))
		}
		break
	default:
		fmt.Println(fmt.Sprintf("has type \"%s\" with value [%v]", reflect.TypeOf(in).Name(), reflect.ValueOf(in)))
	}
}

func main() {
	inspect("This is a string")
	inspect(12345)
	inspect(1.2345)
	inspect(true)
	inspect(Customer{
		FName:        "John",
		LName:        "Wick",
		UserID:       123123123,
		InvoiceTotal: 10000,
	})
}
