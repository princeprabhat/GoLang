package main


/*Write a function that takes a slice of strings and returns a new slice with all strings that are 
palindromes (i.e., read the same backwards as forwards). 
Input: ["racecar", "level", "hello"] Output: ["racecar", "level"]*/


func PalindromeString(s []string)[]string {
	result:= []string{}
	match := ""
	for _,v := range s{
		word := string(v)
		match = ReverseWord(word) // calling the reverse function will return a reverse string
		if(word == match){
			result = append(result,word)
			match = ""
			word = ""
		}

	}
	return result
}


//Function returns a reverse word
func ReverseWord(s string) string{
	var value string
	for i:=len(s)-1;i>=0;i--{
		value += string(s[i])
	}
	return value

}