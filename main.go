package main

import (
	"fmt"
)

// var GoLangLearning = "Learning GoLang step by step"
// var taskItems = []string{GoLangLearning, "Welcome to our Todo app!", "Watching go crash course!", "Next is GoLang full course!"}

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
	maxItem := 45
	// fmt.Println("ToDo App:", taskItems)

	printTasks(taskItems, maxItem)

	println("----------------")
	taskItems = addTask(taskItems, "Go for web dev")
	taskItems = addTask(taskItems, "Go for web dev with react app")

	println("----------------")

	printTasks(taskItems, maxItem)

}

func printTasks(taskItems []string, itemLimit int) {
	fmt.Println("List of my Tasks")

	for index, task := range taskItems {
		// fmt.Println(index+1, ".", task)
		fmt.Printf("%d. %s\n", index+1, task)
	}

	fmt.Printf("Max items are: %v\n", itemLimit)
}

func addTask(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)
	// printTasks(updatedTaskItems, 23)
	return updatedTaskItems

}
