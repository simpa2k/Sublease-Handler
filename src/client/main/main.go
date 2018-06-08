package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	resp, err := http.Get("http://localhost:3000/apartment/1")
	if err != nil {
		panic("Could not make request to apartment endpoint.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(body)
}
