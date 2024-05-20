package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue = "Hello world"
	PrintMe(printValue)

	var num1 int = 12
	var num2 int = 2
	var result, remainder, err = IntDivision(num1, num2)
	/*
		if err != nil {
		fmt.Printf(err.Error())
		} else if remainder == 0 {
			fmt.Printf("Result: %v", result)
		} else {
			fmt.Printf("Result: %v\nRemainder: %v\n", result, remainder)
		}
	*/
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("Result: %v\n", result)
	default:
		fmt.Printf("Result: %v\nRemainder: %v\n", result, remainder)
	}
}

func PrintMe(printValue string) {
	fmt.Println(printValue)
}

func IntDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
