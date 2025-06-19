package main

import "fmt"

func main() {

	// Arrays are used when you know the size of the array by which you can optimized it. If not then we use SLICES which allocate dynamically

	// numbered sequence of equal lenght
	var nums [5]int

	// array length
	fmt.Println(len(nums))

	nums[2] = 1
	fmt.Println(nums)

	// 2D array
	sums := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(sums)

	

}
