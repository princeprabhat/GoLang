package main

/*Write a function that takes a string and returns a map with the frequency count of each character
 in the string. Input: "hello" Output: {"h": 1, "e": 1, "l": 2, "o": 1}*/

 func CountChar(s string) map[string]int{
	m := make(map[string]int)

	for _, v := range s{
		val,ok := m[string(v)]

		if ok{
			m[string(v)] = val+1
		}else {
				m[string(v)] = 1
				}
		
	}
	return m
 }


