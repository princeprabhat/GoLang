package main

/*Write a function that takes a slice of integers and a number and returns the number
 of times the number appears in the slice. Input: ([1, 2, 3, 2, 4, 5], 2) Output: 2*/

func OccurenceCount(n []int, a int) int{
	var count int
	for _,v := range n{
		if v == a{
			count++
			
		}
	}
	return count

}