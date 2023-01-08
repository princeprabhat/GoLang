package main

/*Write a function that takes a slice of integers and returns the index of the first occurrence 
of a given number. Input: ([1, 2, 3, 2, 4, 5], 2) Output: 1*/

	func IndexOccurences(n []int,a int) int{
		var index int
		for i,v := range n{
			if v == a{
				index = i
				break
			}
		}
		return index
	}