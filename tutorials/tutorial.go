package tutorials

// import "fmt"

// func main() {
// 	fmt.Println("Welcome to Quiz game! ")

// 	var name string

// 	fmt.Printf("Enter your name: ")
// 	fmt.Scan(&name)
// 	fmt.Println(name)

// 	var age uint //unasigne int cant be negative like age else have int , float64

// 	fmt.Printf("Enter your age: ")
// 	fmt.Scan(&age)

// 	fmt.Println(age >= 10)
// 	// var isEmployed = true

// 	position := "Full Stack"
// 	// fmt.Println(name)
// 	// fmt.Println(age)
// 	// fmt.Println(isEmployed)
// 	// fmt.Println(position)

// 	// position = "Go-Lang"
// 	if age >= 10 {
// 		fmt.Printf("Hi %v, Lets start becoming %v dev \n", name, position)

// 		fmt.Printf("MERN stand  for?")
// 		var ans string
// 		var ans2 string
// 		fmt.Scan(&ans, &ans2)

// 		if ans+" "+ans2 == "Mongo Express" {
// 			fmt.Println("Correct ")
// 			fmt.Println(ans, ans2)
// 		} else {
// 			fmt.Println("INcorrect")
// 		}

// 	} else {
// 		fmt.Println("Not old yet")
// 		return
// 	}

// 	// fmt.Println("cant play")
// }

// Arrary have fix size and slices is unlimited array
// var taskItems = []string{GoLangLearning, "1. Welcome to our Todo app!", "2. Watching go crash course!", "3. Next is GoLang full course!"}

// fmt.Println("ToDo App:", taskItems)

// for loop syntex
// for index, task := range taskItems {
// 	fmt.Println(index, task)
//  fmt.Printf("%d. %s\n", index+1, task)

// var taskItems = []string{GoLangLearning, "Welcome to our Todo app!", "Watching go crash course!", "Next is GoLang full course!"}
// 	maxItem := 45
// 	// fmt.Println("ToDo App:", taskItems)

// 	printTasks(taskItems, maxItem)

// 	println("----------------")
// 	taskItems = addTask(taskItems, "Go for web dev")
// 	taskItems = addTask(taskItems, "Go for web dev with react app")

// 	println("----------------")

// 	printTasks(taskItems, maxItem)

// }

// func printTasks(taskItems []string, itemLimit int) {
// 	fmt.Println("List of my Tasks")

// 	for index, task := range taskItems {
// 		// fmt.Println(index+1, ".", task)
// 		fmt.Printf("%d. %s\n", index+1, task)
// 	}

// 	fmt.Printf("Max items are: %v\n", itemLimit)
// }

// func addTask(taskItems []string, newTask string) []string {
// 	var updatedTaskItems = append(taskItems, newTask)
// 	// printTasks(updatedTaskItems, 23)
// 	return updatedTaskItems

// }