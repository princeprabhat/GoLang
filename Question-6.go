package main

/*Write a function that takes a slice of strings and returns a new slice with all strings
that have more than 5 characters. Input: ["cat", "dog", "elephant", "lion"] Output: ["elephant"]*/

func CharCount(s []string) []string{
	res := []string{}

	for _, v := range s{
		if len(v)>5{
			res = append(res,v)
		}
	}
	return res
}