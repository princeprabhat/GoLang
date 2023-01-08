package main

/*Write a function that takes a slice of integers and returns true if the slice contains a given number. 
Input: ([1, 2, 3, 4, 5], 3) Output: true*/

func MatchNumSlice(n []int, a int) bool{
	for _,v := range n{
		if v == a{
			return true
		}
	}
	return false
}