package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// write goroutine with different variants for function call.

	// goroutine function call
	go fun("goroutine func call")

	// goroutine with anonymous function
	go func() {
		fun("goroutine-anonymous-func")
	}()

	// goroutine with function value call
	funcValue := fun
	go funcValue("goroutine-with-function-value-call")

	// wait for goroutines to end
	fmt.Println("waiting for go routine...")
	time.Sleep(100 * time.Millisecond) // need this bcoz main routine is not waiting for the the go routine to finish as of now

	fmt.Println("done..")
}
