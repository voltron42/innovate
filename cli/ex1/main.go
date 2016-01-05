package main

import (
	"fmt"
	"os"
)

func main() {
	name := "World"
	args := os.Args
	if len(args) >= 2 {
		name = args[1]
	}
	fmt.Printf("Hello %v!\n", name)
}
