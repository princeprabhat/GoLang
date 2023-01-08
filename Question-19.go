package main
import (
    
    "strings"
)
/*Write a function that takes a string and returns a new string with all the words in reverse order. 
Input: "This is a test" Output: "test a is This"*/


func ReverseString(s string)string {
var result string
words:= strings.Split(s," ")

	for i:= len(words)-1;i>=0;i--{
		result += words[i]
		if i!=0{
			result += " "
		}
	}
	return result
}