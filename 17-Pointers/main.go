package main

import "fmt"

func zeroVal(ival int) {
	ival = 0
}

func zeroPtr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("Initial:", i)

	zeroVal(i)
	fmt.Println("zeroVal:", i)

	fmt.Println("&i mean: ", &i)
	ptr := &i
	fmt.Println("*i mean: ", *ptr)

	zeroPtr(&i)
	fmt.Println("zeroPtr:", i)

}
