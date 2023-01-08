package main

/*Write a function that takes a slice of integers and returns true if the slice is sorted in descending order.
Input: [5, 4, 3, 2, 1] Output: true*/

func IsSortedDesc(n []int) bool{
	res := false
	flag := 0

	for i:=1;i<len(n)-1;i++{
		if n[i] > n[i-1]{
			flag = 1
			
		}
	}
	if flag==0{
		res = true

	}
	return res
}