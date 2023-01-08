package main

/*Write a function that takes a slice of integers and returns a new slice with the elements in reverse order.
 Input: [1, 2, 3, 4, 5] Output: [5, 4, 3, 2, 1]*/

 func ReverseSlice(n []int) []int{
	
	newSlice := []int{}

	for i:= len(n)-1; i>=0; i-- {
		newSlice = append(newSlice, n[i])
	}
	return newSlice
 }