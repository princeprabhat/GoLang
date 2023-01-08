package main
/*Write a function that takes a slice of strings and returns a new slice with all strings 
that start with a given letter. Input: (["cat", "dog", "elephant"], "e") Output: ["elephant"]*/

func StartsWithLetter(s []string, letter string) []string{
	var newSlice []string

	for _, v := range s{
		for _,l := range v{
			if letter == string(l){
				newSlice = append(newSlice, v)
				break
			}
			break
		}
	}
	return newSlice
}