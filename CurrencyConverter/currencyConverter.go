package main

import (
	"errors"
	"fmt"
	"strings"
)

type Currency struct {
	Name         string
	ExchangeRate float64
}

type Converter struct {
	currencyMap map[string]*Currency
}

func (c *Converter) AddCurrency(code string, currency *Currency) {
	if c.currencyMap == nil {
		c.currencyMap = make(map[string]*Currency, 0)
	}

	c.currencyMap[code] = &Currency{
		Name:         currency.Name,
		ExchangeRate: currency.ExchangeRate,
	}
}

func (c *Converter) PrintCurrency(code string) {
	if v, ok := c.currencyMap[code]; ok {
		fmt.Println(code, " | ", v.Name, "| ", v.ExchangeRate)
	} else {
		fmt.Println(fmt.Sprintf("currency %s does not exist", code))
	}
}

func (c *Converter) PrintCurrencies() {
	for k, v := range c.currencyMap { // TODO: format this to align | between rows
		fmt.Println(k, " | ", v.Name, "| ", v.ExchangeRate)
	}
}

func (c *Converter) Convert(amount float64, from, to string) (float64, error) {
	res := 0.0
	var err error
	err = nil

	//no conversion required
	if from == to {
		return amount, nil
	}

	if fromCurr, ok := c.currencyMap[from]; ok {
		if toCurr, ok := c.currencyMap[to]; ok {
			if fromCurr.ExchangeRate == 0.0 { // TODO: might want to consider epsilon instead of 0.0
				err = errors.New("division by 0")
			} else {
				res = amount / fromCurr.ExchangeRate * toCurr.ExchangeRate
			}
		} else {
			err = errors.New(fmt.Sprintf("currency %s does not exist", to))
		}
	} else {
		err = errors.New(fmt.Sprintf("currency %s does not exist", from))
	}

	return res, err
}

func getInputCurrency(prompt string) string {
	in := ""

	for {
		fmt.Printf("%s: ", prompt)
		_, err := fmt.Scan(&in)
		fmt.Println()

		if err != nil {
			fmt.Println("Error processing input: ", err)
			continue
		}

		break
	}

	return in
}

func getInputAmount() float64 {
	in := 0.0

	for {
		fmt.Printf("Enter amount to convert: ")
		_, err := fmt.Scan(&in)
		fmt.Println()

		if err != nil {
			fmt.Println("Error processing input: ", err)
			continue
		}

		break
	}

	return in
}

func main() {
	converter := &Converter{}
	initCurrencies(converter)

	from := getInputCurrency("Please enter currency to convert from")
	amount := getInputAmount()
	to := getInputCurrency("Please enter currency to convert to")

	from = strings.ToUpper(from)
	to = strings.ToUpper(to)

	res, err := converter.Convert(amount, from, to)

	if err != nil {
		fmt.Println("Encountered error: ", err)
	} else {
		// assume 2 dec place for all currencies
		fmt.Println(fmt.Sprintf("Converted %.2f%s into %.2f%s", amount, from, res, to))
	}
}

func initCurrencies(c *Converter) {
	c.AddCurrency("USD", &Currency{
		Name:         "US dollar",
		ExchangeRate: 1.1318,
	})

	c.AddCurrency("JPY", &Currency{
		Name:         "Japanese yen",
		ExchangeRate: 121.05,
	})

	c.AddCurrency("GBP", &Currency{
		Name:         "Pound sterling",
		ExchangeRate: 0.90630,
	})

	c.AddCurrency("CNY", &Currency{
		Name:         "Chinese yuan renminbi",
		ExchangeRate: 7.9944,
	})

	c.AddCurrency("SGD", &Currency{
		Name:         "Singapore dollar",
		ExchangeRate: 1.5743,
	})

	c.AddCurrency("MYR", &Currency{
		Name:         "Malaysian ringgit",
		ExchangeRate: 4.8390,
	})
}
