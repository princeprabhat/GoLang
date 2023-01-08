package main

/*Write a function that takes a slice of strings and returns a new slice with all strings that
are at least a given length. Input: (["cat", "dog", "elephant", "lion"], 5) Output: ["elephant"]*/

func WordLength(s []string, n int) []string{
	result := []string{}

	for _,v := range s{
		if len(v) >= n{
			result = append(result,v)
		}
	}
	return result
}