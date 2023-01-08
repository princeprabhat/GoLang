package main

/*Write a function that takes a string and returns a new string with all vowels removed.
 Input: "hello" Output: "hll"*/

func RemoveVowels(s string) string{
	var a string 

	for _,v := range s{
		if v=='a' || v=='e' || v=='i' || v=='o' || v=='u' || v=='A' || v=='E' || v=='I' || v=='O' || v=='U'{
			continue
		}else{
			a = a+string(v)
		}


	}
	return a
}