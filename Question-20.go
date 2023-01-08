package main

/*Write a function that takes a slice of integers and returns a new slice with all the prime numbers.
 Input: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] Output: [2, 3, 5, 7]*/

func IsPrime(n []int) []int{
	result := []int{}

	for _,v := range n{
		if checkPrime(v){
			result = append(result,v)
		}
	}
	return result
}

//Function to check if a number is prime

func checkPrime(a int) bool{
	count :=0
	if a ==1{
		return false
	}
	for i:=2;i<=a;i++{
		if a%i ==0{
			count++
		}
		if count>1{
			break
		}
	}
	if count == 1{
		return true
	}
	return false
}