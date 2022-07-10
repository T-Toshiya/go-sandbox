package main

import (
	"fmt"
	"sort"
)

func main() {
	target := 39
	data := []int{3, 5, 8, 10, 14, 17, 21, 39}

	i := sort.Search(len(data), func(i int) bool { return data[i] >= target })
	if i < len(data) && data[i] == target {
		fmt.Printf("%d is present at data[%d]", target, i)
	} else {
		fmt.Printf("%d is not present in data", target)
	}
}
