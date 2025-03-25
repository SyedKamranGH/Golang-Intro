package tutorials

// import (
// 	"fmt"
// 	"net/http"
// )

// var taskItems = []string{"Welcome to our Todo app!", "Watching go crash course!", "Next is GoLang full course!"}

// func main() {
// 	fmt.Println("############ Welcome to our Todo List App! ########")

// 	http.HandleFunc("/", helloUser)
// 	http.HandleFunc("/show-tasks", showTasks)
// 	http.ListenAndServe(":8080", nil)
// }

// func helloUser(w http.ResponseWriter, req *http.Request) {
// 	greeting := "Help user. Welcome to our Golang TodoList App"
// 	fmt.Fprintf(w, greeting)
// }

// func showTasks(w http.ResponseWriter, req *http.Request) {
// 	for index, task := range taskItems {

// 		fmt.Fprintln(w, index, task)
// 	}

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
