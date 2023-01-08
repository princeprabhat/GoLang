package main

/*Write a function that takes a slice of integers and returns the average of all the elements in the slice.
 Input: [1, 2, 3, 4, 5] Output: 3*/

 func AvgSlice(a []int) int{
	var sum int
	for _,v := range a{
		sum +=v
	}
	res := sum/len(a)
	return res
 }