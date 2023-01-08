package main
import "sort"
/*Write a function that takes a slice of integers and returns the second-largest number in the slice. 
Input: [1, 2, 3, 4, 5] Output: 4 */

func SecondLargestNumber(a []int) int{
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
	sort.Ints(res)

	return res[len(res)-2]
}