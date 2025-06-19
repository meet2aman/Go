package main

import (
	"fmt"
	"time" // its a time package with gives current time, weekday, year, month, etc.
)

func main() {
	// simple switch
	i := 1
	switch i {
	case 1:
		fmt.Println("One")
		// break // we dont need to write break specifically Its LEXER detect automatically

	case 2:
		fmt.Println("two")

	case 3:
		fmt.Println("Three")

	case 4:
		fmt.Println("Four")

	default:
		fmt.Println("None")

	}

	/// With Multiple Condition Switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's a weekend")
	default:
		fmt.Println("It's a workday")
	}

	/// Type Switch
	whoAmI := func(i interface{}) { /// interface{} is a of any type any OR We can also write it as [any]
		switch t := i.(type) {
		case int:
			fmt.Println("It's an Integer")
		case string:
			fmt.Println("It's a string")
		case bool:
			fmt.Println("It's a boolean")
		case float32:
			fmt.Println("It's a Float")
		default:
			fmt.Println("Other", t)
		}

	}

	whoAmI(true)

}
