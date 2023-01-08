package main

/*Write a function that takes a slice of integers and returns true if the slice is sorted in ascending order. 
Input: [1, 2, 3, 4, 5] Output: true*/

func IsSorted(n []int) bool{
	res := false
	flag := 0

	for i:=1;i<len(n)-1;i++{
		if n[i] < n[i-1]{
			flag = 1
			
		}
	}
	if flag==0{
		res = true

	}
	return res
}