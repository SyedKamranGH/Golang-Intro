package main

import "fmt"

type person struct {
	name string
	age  int
	sex  string
}

func newPerson(name string, sex string) *person {

	p := person{name: name, age: 42, sex: sex}
	// p.age = 42
	// p.sex = sex
	// p.sex = person{sex: sex}
	return &p
}

func main() {

	// address := 1330

	fmt.Println(person{"Bob", 20, "Male"})

	fmt.Println(person{name: "Alice", age: 30, sex: "Female"})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", sex: "Female"})

	fmt.Println(newPerson("Jon", "Male"))

	s := person{name: "Sean", age: 50, sex: "xyz"}
	fmt.Println(s, s.name, s.age)

	sp := &s
	sp.age = 79
	fmt.Println(sp.age)
	fmt.Println(s)

	// sp.age = 51
	// fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

	// fmt.Println(dog{name: "Scooby", isGood: "true"})
}
