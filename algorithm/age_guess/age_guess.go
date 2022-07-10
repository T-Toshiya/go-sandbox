package main

import "fmt"

func main() {
	fmt.Println("Start Game!")

	left := 20
	right := 36

	for right-left > 1 {
		mid := left + (right-left)/2

		fmt.Printf("Is the age less than %d ?(yes / no) \n", mid)

		var ans string
		_, err := fmt.Scan(&ans)
		if err != nil {
			return
		}

		if ans == "yes" {
			right = mid
		} else {
			left = mid
		}
	}

	fmt.Printf("The age is %d!", left)
}
