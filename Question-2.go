package main

import "sort"

/*Write a function that takes a slice of strings and returns a new slice with all strings 
sorted in alphabetical order. 
Input: ["zebra", "monkey", "aardvark"] Output: ["aardvark", "monkey", "zebra"]*/

func SortSlice(s []string) []string{
	sort.Strings(s)
	return s
}