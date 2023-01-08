package main
import (
	
	"strings"
)
/*Write a function that takes a string and returns a new string with all the words in reverse order,
 with the letters in each word reversed as well. Input: "This is a test" Output: "tset a si sihT"*/

	


func ReverseStringWords(s string) string{
	var result string

	words:=ReverseString(s) // Using our own reverse string function
	words1 := strings.Split(words, " ")
	for _, v := range words1{
		result +=ReverseWord(string(v)) + " " // Using our own reverse word function
	}
	
	return result
}