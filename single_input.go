package main

import "fmt"

func main() {
	var name string
	fmt.Print("enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("hey there, ", name)

}
