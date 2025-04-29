package main

import (
	"fmt"
)

func main() {
	// this comment doesnt add any value and should trigger a review comment instead..
	// fmt.Prinln("hello world")

	x := 42
	if x == 42 {
		fmt.Println("The answer is correct.")
	}

	// Inefficient string concatenation in a loop
	result := ""
	for i := 0; i < 5; i++ {
		result = result + fmt.Sprintf("Item %d, ", i)
	}
	fmt.Println(result)

	fmt.Println("echo hello world")
}
