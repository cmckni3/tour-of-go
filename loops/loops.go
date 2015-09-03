package main

import(
  "fmt"
)

func main() {
  list := []string { "one", "two", "three" }

  for index, val := range list {
    fmt.Println(index, val)
  }
  for i:=0; i<20; i++ {
    fmt.Print("Go! ")
  }
}
