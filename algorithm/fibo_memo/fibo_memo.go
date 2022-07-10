package main

import "fmt"

func fiboMemo(n int, memo []int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	if memo[n-1] != 0 {
		return memo[n-1]
	}

	memo[n-1] = fiboMemo(n-1, memo) + fiboMemo(n-2, memo)

	return memo[n-1]
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	memo := make([]int, n)

	fiboMemo(n, memo)

	for i, v := range memo {
		fmt.Printf("%d 項目: %d \n", i+1, v)
	}
}
