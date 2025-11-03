package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("sending", i)
			c <- i
		}
	}()
	go func() {
		for {
			fmt.Println("receiving", <-c)
		}
	}()
	time.Sleep(1 * time.Second)
}
