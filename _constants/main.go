package main

import "fmt"

// These declaration are GLOBAL available
const age int = 30 // you can declare this const out side the function block
var language string = "golang"

// username := "Meet2Aman" /// We can't do this out the scope or function Block

func main() {

	/// WallRuss Operator
	username := "Meet2Aman" // This Is Shorthand Notation in Which username treated as VAR not CONST and it automatically infer the type also

	const Fname string = "Thor" // We Can declare this as well
	const Lname = "Loki"        //  or this as well It Automtically infer The Type of it Using LEXER
	fmt.Println(Fname, Lname, username, language, age)

	// GROUP DECLARATION
	const (
		port = 8080
		host = "localhost"
	)

	fmt.Println(`The Go is running on`, host, ":", port)

}
