package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 5; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(3000)
	}
}

func main() {

	go f("direct")
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("anonimous")

	//fmt.Scanln()
	fmt.Println("done")
}
