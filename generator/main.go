package main

import "fmt"

//go:generate go run main.go generator.go
func main() {
	if err := Generate(); err != nil {
		fmt.Println("error: ", err)
	}
}
