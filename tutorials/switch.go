package tutorials

import (
	"fmt"
	"time"
)

func checkSaturday() {
	fmt.Println("When Saturday is coming")
	today := time.Now().Weekday()
	fmt.Println("Today: ", today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")

	}
}

func dailtGreetings() {
	t := time.Now()
	fmt.Println("Current Time: ", t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening")
	}
}

// func main() {
// 	checkSaturday()
// 	dailtGreetings()
// }
