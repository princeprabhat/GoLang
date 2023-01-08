package main
//Write a function that takes a slice of integers and returns a new slice with only the even numbers. 
//Input: [1, 2, 3, 4, 5] Output: [2, 4]
func EvenSlice(a []int) []int{
	res := []int{}
	for _,v := range a{
	if v%2==0{
		res=append(res,v)
	}
		
	}
	return res
}
