package main

import (
	"fmt"
	"strconv"
)

type channel chan int

func generate(c channel) {
	defer fmt.Println("Generation ended")
	for i := 0; i < 3; i++ {
		c <- i
		fmt.Println(strconv.Itoa(i) + " generated")
	}
	close(c)
}
func pass(c1, c2 channel) {
	defer fmt.Println("Passing ended")
	for val := range c1 {
		c2 <- val
		fmt.Println(strconv.Itoa(val) + " passed")
	}
	close(c2)
}
func main() {
	defer fmt.Println("Bye")
	ch1 := make(channel, 3)
	ch2 := make(channel, 3)
	go generate(ch1)
	go pass(ch1, ch2)
	for a := range ch2 {
		fmt.Println(strconv.Itoa(a) + " read")
	}
	fmt.Scanln()
}
