package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getLuckyNum(c chan<- int) {
	fmt.Println("...")

	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

	num := rand.Intn(10)
	c <- num
}

func iterate() {
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func main() {
	iterate()
	fmt.Println("what is today's lucky number?")

	c := make(chan int)
	go getLuckyNum(c)

	num := <-c

	fmt.Printf("Today's your lucky number is %d!\n", num)

	close(c)
}
