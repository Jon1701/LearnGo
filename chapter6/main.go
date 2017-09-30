package main

import "fmt"

func main() {
  // Declare array of 5 ints (default to 0 each)
  var x [5]int

  // Set 5th element to 100.
  x[4] = 100

  fmt.Println(x)

  // Declare array of 5 float64.
  var y [5]float64

  // Set elements.
  y[0] = 98
  y[1] = 93
  y[2] = 77
  y[3] = 82
  y[4] = 83

  // Create a variable of type float64.
  var total float64 = 0

  // Calculate total.
  for _, value := range y {
    total += value
  }

  // Print average.
  fmt.Println(total / float64(len(y)))

  // Shorter syntax for creating arrays:
  z := [5]float64{ 98, 93, 77, 82, 83 }
  fmt.Println(z)
}
