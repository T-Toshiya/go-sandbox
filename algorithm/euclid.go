package main

import "fmt"

func euclid(m, n int) int {
	r := m % n
	if r == 0 {
		return n
	}

	return euclid(n, r)
}

func main() {
	var m, n int
	_, err := fmt.Scan(&m)
	if err != nil {
		return
	}
	_, err = fmt.Scan(&n)
	if err != nil {
		return
	}

	fmt.Println(euclid(m, n))
}
