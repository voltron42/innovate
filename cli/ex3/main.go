package main

import "fmt"

func main() {
	fmt.Println("")
	fmt.Println("counting")

	defer fmt.Println("done")

	for i := 0; i < 5; i++ {
		fmt.Printf("Deferring %v\n", i)
		defer fmt.Println(i)
	}

	fmt.Println("")
	fmt.Println("start countdown")
}
