package main

/*Write a function that takes a slice of strings and a string, and returns the index of the first
occurrence of the string in the slice. Input: (["cat", "dog", "elephant"], "dog") Output: 1*/

func StringMatch(s []string, a string) int{
	output := -1
	for i,v := range s {
		if v == string(a){
			output = i
			break
		}
	}
	return output
}