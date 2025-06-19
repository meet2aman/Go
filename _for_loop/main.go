package main

import "fmt"

// GO HAS ONLY FOR LOOP CONSTRUCT != NO WHILE OR DO WHILE LOOP BUT WE CAN MIMIC THEM USING FOR LOOP

func main() {
	// While loop using FOR LOOP
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// infinte Loop
	for {
		fmt.Println("∞")
		break
	}

	/// Classic FOR LOOP
	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			if i > 10 {
				break // break the loop
			}
		}
		continue // pass the loop 
	}

	/// Go 1.22 version there is RANGE
	for i := range 3 {
		fmt.Println(i)
	}
}
