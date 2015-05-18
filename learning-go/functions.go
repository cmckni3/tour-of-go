package main

import "fmt"

// equivalent to func add(x int, y int) int {
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

