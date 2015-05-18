package main

import(
  "fmt"
  "strings"
)

func main() {
  fmt.Println(double(5))


  first, last := parseName("John Doe")
  fmt.Println(first)
  fmt.Println(last)
}

func double(n int) int {
  return n + n
}

func parseName(name string) (first, last string) {
  parsed := strings.Split(name, " ")
  return parsed[0], parsed[1]
}

