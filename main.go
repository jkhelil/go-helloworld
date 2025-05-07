package main

import (
	"fmt"
)

func main() {
	// this is no purpose comment and should trigger a review comment
	// fmt.Prinln("hello world")

	// new comment

	// comment bizare

	x := 40
	if x == 40 {
		fmt.Println("x is 40")
	} else {
		fmt.Println("x is not 40")
	}

	secret := "my secret"
	if secret == "my secret" {
		fmt.Println("secret is my secret")
	} else {
		fmt.Println("secret is not my secret")
	}

	fmt.Println("echo hello world")
}
