package main

/*Write a function that takes a slice of integers and returns a new slice with all elements that 
are palindromes (i.e., read the same forwards as backwards). 
Input: [1, 11, 121, 12321] Output: [11, 121, 12321]*/

func PalindromeIntegers(n []int) []int{
	result := []int{}
	var check int
	var remainder int
	for _,v:= range n{
		if v <=9 {
			continue
		}
		temp := v
		for temp!=0{
			remainder = temp % 10;
			check = check * 10 + remainder;
			temp /= 10;
		}
		if check == v{
			result = append(result,v)
		}
		remainder = 0
		check = 0
	}
	return result

}