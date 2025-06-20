package main

import (
	"fmt"
	"slices"
)

// slices -> dynamic
// most used construct in go
// + useful methods

func main() {
	// un-utilised slices is nil (null)
	var nums []int
	fmt.Println(nums)      // output []
	fmt.Println(len(nums)) // output 0

	var sums = make([]int, 2, 5)
	fmt.Println(sums) // output is not nill -> [0,0]
	// 2 is where we give capacity of slice although it can shrink or grow. it's the maximum length

	fmt.Println(cap(sums)) // cap method check the capacity of slice

	sums = append(sums, 1) // insert the value in the slice from end
	fmt.Println(sums)      // output -> [0 0 1]

	// ! if we do more appends and it exceed more then 5 then capacity doubles from its present capacity
	fmt.Println(cap(sums)) // Output -> capacity is 5. Then output capacity becomes 10

	/// SLICE OPERATOR
	var nums1 = []int{1, 2, 3,5}
	fmt.Println(nums1[0:1]) // output -> [1] from 0 index to 1st index but exclude the 1st index
	fmt.Println(nums1[:2])  // output -> [1 2] from 0 index to 2nd index but exclude the 2nd index
	fmt.Println(nums1[0:])  // output -> [1 2 3] from 0 index to last index availabe

	fmt.Println(slices.Equal(nums, nums1)) // output -> false | There is slices package available which has a Equal method to check two slices are equal or not

	/// 2D slices
	var num3 = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(num3) // output -> [1 2 3][4 5 6]

	/// copy elements from one arrays to another 
	var nums2 = make([]int, len(nums1))
	copy(nums2, nums1)
	fmt.Println(nums2)

}
