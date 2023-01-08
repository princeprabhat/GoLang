package main

/*Write a function that takes a slice of integers and returns a new slice with all duplicates removed. 
Input: [1, 2, 3, 2, 4, 5, 5] Output: [1, 3, 4]*/

func DuplicateSlice(a []int)[]int{
	 res := []int{}
	 flag:=0
	for i:=0;i<len(a);i++{
		flag=0
		for j:=0;j<i;j++{
			if a[i]==a[j]{
				flag=1
			}

		}
		if flag ==0 {
			res = append(res,a[i])
			
		}
	}
	 return res
}