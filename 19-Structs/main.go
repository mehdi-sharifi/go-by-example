package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	p := person{name: name, age: age}
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "John", age: 24})
	fmt.Println(person{name: "Jeff"})
	fmt.Println(&person{name: "Adam", age: 44})
	fmt.Println(newPerson("Mike", 47))

	s := person{name: "Lee", age: 29}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 68
	fmt.Println(sp.age)
	fmt.Println(s.age)


	dog := struct {
		name string
		isGood bool
	}{
		"Alex",
		true,
	}

	fmt.Println(dog)

}
