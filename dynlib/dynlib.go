package main

import "C"
import "fmt"

//export Print
func Print(message string) {
	fmt.Println(message)
}

func main() {}
