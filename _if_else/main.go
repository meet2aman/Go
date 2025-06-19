package main

import "fmt"

func main() {
	age := 60

	// if else
	if age >= 18 {
		fmt.Println("Your can drive a car")
	} else {
		fmt.Println("Your can't drive a car")
	}

	/// else if
	if age >= 60 {
		fmt.Println("Your are Old")
	} else if age >= 18 {
		fmt.Println("Your are Adult")
	} else if age > 12 {
		fmt.Println("Your are Teenager")
	} else if age < 12 {
		fmt.Println("Your are Kid")
	}
}
