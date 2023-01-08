package main

/*Write a function that takes a slice of strings and returns a new slice with all strings that 
contain a given letter. Input: (["cat", "dog", "elephant", "lion"], "e") Output: ["elephant"]*/

func LetterString(s []string, a string) []string{
	result := []string{}

	for _,v := range s{
		for _,v2 := range v{
			if string(v2) == a{
				result = append(result, v)
				break
			}
		}
	}
	return result
}