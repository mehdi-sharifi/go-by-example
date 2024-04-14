package main

import "fmt"

func mathematical(a, b float64) (plus float64, minus float64, multiplication float64, division float64) {
	plus = a + b
	minus = a - b
	multiplication = a * b
	division = a / b
	return
}

func main() {
	a := 14.0
	b := 4.0

	plusVal, minusVal, multiVal, divVal := mathematical(a, b)

	fmt.Printf("%f + %f = %f\n", a, b, plusVal)
	fmt.Printf("%f - %f = %f\n", a, b, minusVal)
	fmt.Printf("%f * %f = %f\n", a, b, multiVal)
	fmt.Printf("%f / %f = %f\n", a, b, divVal)

}
