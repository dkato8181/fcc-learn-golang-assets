package main

import "fmt"

func main() {
	var username string = "wagslane"
	var password string = "20947382822"

	fmt.Println("Authorization: Basic", username+":"+password)
}
