package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Api Application")

	response, err := http.Get("https://google.com")

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	str, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(str))

}
