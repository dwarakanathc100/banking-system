package main

import "fmt"

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("First loop", ":", i)
	}
}

func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("Second loop :", i)
	}
}

func main() {
	test1()
	test2()
}
