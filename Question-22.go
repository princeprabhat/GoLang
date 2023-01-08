package main

/*Write a function that takes a slice of integers and returns a new slice with all elements 
that are divisible by a given number. Input: ([1, 2, 3, 4, 5], 2) Output: [2, 4]*/

func DivisibleByNumber(n []int, a int) []int{
	result := []int{}
	for _,v := range n{
		if v%a==0{
			result = append(result,v)
		}
	}
	return result
}