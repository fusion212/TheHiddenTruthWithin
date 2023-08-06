package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Hello humans. This is my little request :)")
	Eldritch_Wand()
	PerformGetRequest()
}

func Eldritch_Wand() {
	const Julia = "Christopher... Can you see the truth?"
	fmt.Println(Julia)

	for i := 1; i <= 60; i++ {
		fmt.Println("Attempt:", i, Julia)
	}
}

func PerformGetRequest() {
	const Frank = "http://127.0.0.1:8000/get"
	response, err := http.Get(Frank)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	ContentLength, _ := io.ReadAll(response.Body)
	fmt.Println(string(ContentLength))

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)
	fmt.Println("User-Agent:", response.Request.UserAgent())
	fmt.Println("Frank, is", Frank, "your address?")
}
