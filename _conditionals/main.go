package main

import "fmt"

func main() {
	var role = "admin"
	var hasPermission = true

	if role == "admin" || hasPermission { // OR Operator if One condition is true block will execute.
		fmt.Println("You are Authorised ✅")
	} else {
		fmt.Println("You are not Authorised ❌")
	}

	if role == "admin" && hasPermission { // AND Operator When BOTH conditions are true then block will execute.
		fmt.Println("You are Authorised ✅")
	} else {
		fmt.Println("You are not Authorised ❌")
	}

	if role := "user"; role == "admin" || hasPermission {
		 // HERE you can directly declare or re-assign value of variable while defining if-else condition
		if role != "admin" { // NOT operator
			fmt.Println("You are denied because You are not a Admin.")
		} else if hasPermission == false {
			fmt.Println("You are denied because You don't have Permissions.")
		} else {
			fmt.Println("You are Allowed")
		}
	} else {
		fmt.Println("You are denied because You don't have Permission and You are not a Admin")

	}

}
