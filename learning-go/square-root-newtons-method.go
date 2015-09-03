package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
  z := 1.0
  delta := 1.0E-6
  previousZ := 0.0
  i := 0
  z = z - ((math.Pow(z, 2) - x) / (2*z))
  for math.Abs(previousZ - z) > delta {
	  previousZ = z
	  z = z - ((math.Pow(z, 2) - x) / (2*z))
	  i = i + 1
  }
  return z
}

func main() {
	fmt.Println(Sqrt(2))
}
