package main

import (
	"../../../yogiberra"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	value, err := retrieveAndValidate(args)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%v is a valid value\n", value)
}

func retrieveAndValidate(args []string) (int, error) {
	if len(args) < 2 {
		panic("Must have a parameter")
	}
	arg1 := args[1]
	value, err := strconv.Atoi(arg1)
	if err != nil {
		return 0, err
	}
	err = yogi.Catch(func() {
		validate(value, 3, true)
	})
	if err != nil {
		return 0, err
	}
	err = validate(value, 5, false)
	return value, err
}

func validate(value, factor int, shouldPanic bool) error {
	if value%factor == 0 {
		return nil
	}
	err := fmt.Errorf("%v is not a factor of %v", factor, value)
	if shouldPanic {
		panic(err)
	} else {
		return err
	}
}
