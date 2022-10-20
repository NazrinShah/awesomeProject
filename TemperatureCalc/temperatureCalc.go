package main

import "fmt"

const (
	OptionKelvin     = 1
	OptionCelsius    = 2
	OptionFahrenheit = 3
)

var opToUnit = map[int]string{
	OptionKelvin:     "K",
	OptionCelsius:    "C",
	OptionFahrenheit: "F",
}

func convertFToK(input float64) float64 {
	return float64(5) / float64(9) * (input + 459.67)
}

func convertFToC(input float64) float64 {
	return float64(5) / float64(9) * (input - 32)
}

func convertKToF(input float64) float64 {
	return float64(9)/float64(5)*(input) - 459.67
}

func convertKToC(input float64) float64 {
	return input - float64(5*32/9) - 459.67
}

func convertCToF(input float64) float64 {
	return float64(9)/float64(5)*input + 32
}

func convertCToK(input float64) float64 {
	return input + float64(5*32/9) + 459.67
}

func getInputUnit() int {
	var sel int

	for {
		fmt.Printf("1: Kelvin\n2: Celsius\n3: Fahrenheit\nSelect input temerature unit: ")
		_, err := fmt.Scan(&sel)
		fmt.Println()

		if err != nil {
			fmt.Println("Error processing input: ", err)
			continue
		} else if sel < OptionKelvin || sel > OptionFahrenheit {
			fmt.Println("Expecting input between 1 to 3 but got ", sel)
			continue
		}

		break
	}

	return sel
}

func getInputTemp() float64 {
	var input float64

	for {
		fmt.Printf("Enter a number to convert: ")
		_, err := fmt.Scan(&input)

		if err != nil {
			fmt.Println("Error processing input: ", err)
			continue
		}

		break
	}

	return input
}

func main() {
	convMap := map[int]map[int]func(float64) float64{
		OptionKelvin: {
			OptionCelsius:    convertKToC,
			OptionFahrenheit: convertKToF,
		},
		OptionCelsius: {
			OptionKelvin:     convertCToK,
			OptionFahrenheit: convertCToF,
		},
		OptionFahrenheit: {
			OptionKelvin:  convertFToK,
			OptionCelsius: convertFToC,
		},
	}

	sel := getInputUnit()
	input := getInputTemp()

	for k, f := range convMap[sel] {
		fmt.Println(fmt.Sprintf("%.2f%s = %.2f%s", input, opToUnit[sel], f(input), opToUnit[k]))
	}
}
