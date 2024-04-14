package main

import "fmt"

func sum(nums ...int) int {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	sumValues := sum(1, 2, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14)
	fmt.Println("Sum Values:", sumValues)

	nums := []int{1, 2, 4, 5, 6, 7, 8}
	sumValues = sum(nums...)
	fmt.Println("Sum Values:", sumValues)

}
