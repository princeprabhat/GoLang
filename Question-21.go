package main

/*Write a function that takes a slice of integers and returns a new slice with the elements in
 alternating order. Input: [1, 2, 3, 4, 5] Output: [1, 3, 5, 2, 4]*/

func AlterNumber(n []int) []int{
	result := []int{}
	for i:=0;i<len(n);i=i+2{
		result = append(result,n[i])
	}
	for i:=1;i<len(n);i=i+2{
		result = append(result,n[i])
	}
	return result

}