package main

import (
	"fmt"
	"reflect"
)

type uid int

type Customer struct {
	FName        string
	LName        string
	UserID       uid
	InvoiceTotal float64
}

func inspect(in interface{}) {
	switch in.(type) {
	case Customer:
		v := reflect.ValueOf(in)
		v1 := reflect.Indirect(v)
		fmt.Println(fmt.Sprintf("Customer has %d fields", v.NumField()))

		for i := 0; i < v1.NumField(); i++ {
			v1Type := v1.Type()
			fmt.Println(fmt.Sprintf("%s is of type %s and %s kind and has value %+v", v1Type.Field(i).Name, v1Type.Field(i).Type.Name(), v.Field(i).Kind(), v.Field(i).Interface()))
		}
	default:
		fmt.Println(fmt.Sprintf("variable has type \"%s\" with value [%v]", reflect.TypeOf(in).Name(), reflect.ValueOf(in)))
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
