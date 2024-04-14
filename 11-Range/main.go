package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("Key:", k)
	}

	for _, v := range kvs {
		fmt.Println("Value:", v)
	}

	str := "hello go!\n"

	for i, c := range str {
		fmt.Println(i, c)
	}

	for i, c := range str {
		fmt.Printf("index:%d --> Unicode:%d --> Character:%c\n", i, c, c)
	}

}
