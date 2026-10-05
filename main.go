package main

import "fmt"

// variadic func
func sumAll(nums ...int) int {
	total := 0

	for _, currVal := range nums {
		total += currVal
	}
	return total
}

func main() {
	// 	views := []int{1, 2, 3, 4, 5, 6}
	//
	// 	// for range
	// 	total := 0
	//
	// 	for i, v := range views {
	// 		fmt.Println("day", i, "views", v)
	// 		total = total + v
	// 	}
	//
	// 	fmt.Println(total)

	// ages := map[string]int{
	// 	"animesh": 21,
	// 	"ujjwal":  20,
	// }
	// fmt.Println(ages, ages["animesh"], len(ages))

	// users := map[string]string{
	// 	"u1": "sangam",
	// 	"u2": "john",
	// 	"u3": "rahul",
	// }
	// fmt.Println(users)

	fmt.Println(sumAll(1, 2, 3, 4, 5))
}
