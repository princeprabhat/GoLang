package main

/*Write a function that takes a slice of strings and a string, 
and returns true if the string is contained in the slice. 
Input: (["cat", "dog", "elephant"], "dog") Output: true*/

func MatchString(s []string, m string) bool{
	var value bool
	for _,v := range s{
		if string(m)==string(v){
			value = true
			break
		}
		
	}
	
	return value

}