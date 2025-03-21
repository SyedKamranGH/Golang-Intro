package main

import (
	"fmt"
)

func main() {
	fmt.Println("~~~~~ Welcome to GoLang! ############")

	GoLangLearning := "Learning GoLang step by step"

	// fmt.Println("1. Welcome to our Todo app!")
	// fmt.Println("2. Watching go crash course!")
	// fmt.Println("3. Next is GoLang full course!")
	// fmt.Println(1231)
	// fmt.Println(GoLangLearning)

	// Arrary have fix size and slices is unlimited array
	var taskItems = []string{GoLangLearning, "Welcome to our Todo app!", "Watching go crash course!", "Next is GoLang full course!"}

	fmt.Println("ToDo App:", taskItems)

	for index, task := range taskItems {
		// fmt.Println(index+1, ".", task)
		fmt.Printf("%d. %s\n", index+1, task)
	}

}
