package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	a := 1
	b := 2
	c := 3
	res := plus(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, res)

	res = plusPlus(a, b, c)
	fmt.Printf("%d + %d + %d = %d\n", a, b, c, res)
}
