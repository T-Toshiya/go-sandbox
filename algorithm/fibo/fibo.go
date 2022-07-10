package main

import "fmt"

func fibo(n int) int {
	fmt.Printf("fibo(%d)を呼び出しました\n", n)

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	result := fibo(n-1) + fibo(n-2)

	fmt.Printf("%d 項目 = %d\n", n, result)

	return result
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	fibo(n)
}
