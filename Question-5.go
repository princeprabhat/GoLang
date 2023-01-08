package main

/*Write a function that takes a slice of integers and returns the largest number in the slice. 
Input: [1, 2, 3, 4, 5] Output: 5*/

func MaxNum(num []int) int{
	res := 0
	for i := 0; i< len(num); i++{
		if num[i] > res{
			res = num[i]
		}
	}
	return res
}