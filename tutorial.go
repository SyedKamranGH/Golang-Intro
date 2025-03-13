package main

import "fmt"

func main() {
	fmt.Println("Welcome to Quiz game! ")

	var name string

	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println(name)

	var age uint //unasigne int cant be negative like age else have int , float64

	fmt.Printf("Enter your age: ")
	fmt.Scan(&age)

	fmt.Println(age >= 10)
	// var isEmployed = true

	position := "Full Stack"
	// fmt.Println(name)
	// fmt.Println(age)
	// fmt.Println(isEmployed)
	// fmt.Println(position)

	// position = "Go-Lang"
	if age >= 10 {
		fmt.Printf("Hi %v, Lets start becoming %v dev \n", name, position)

		fmt.Printf("MERN stand  for?")
		var ans string
		var ans2 string
		fmt.Scan(&ans, &ans2)

		if ans+" "+ans2 == "Mongo Express" {
			fmt.Println("Correct ")
			fmt.Println(ans, ans2)
		} else {
			fmt.Println("INcorrect")
		}

	} else {
		fmt.Println("Not old yet")
		return
	}

	// fmt.Println("cant play")
}
