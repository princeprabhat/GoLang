package main

/*Write a function that takes a slice of integers and returns the smallest number that is divisible
by all elements in the slice. Input: [2, 3, 5] Output: 30*/


func LowestCommonMultiple(s []int) int{
	var result int
	counter :=0
	index :=2
	maxNum:=MaxNum(s) // Calling our MaxNum function will return the maximum number
	temp:= maxNum
for{
	for _,v := range s{
		if temp%v ==0{
			counter++
		}
	}
	

	if counter == len(s){
		result = temp
		break
	}else if counter < len(s){
		counter = 0
	}
	temp=maxNum *index
	index++
}
	return result
}