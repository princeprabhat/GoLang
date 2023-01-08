package main

/*Write a function that takes a slice of integers and returns the sum of all the elements in the slice. 
Input: [1, 2, 3, 4, 5] Output: 15*/

func AddInt(a []int) int{
var result int

for _, v := range a{
	result += v
}
return result

}