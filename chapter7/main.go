package main

import "fmt"

func main() {
  // Create array of float64.
  xs := []float64{
    98,
    93,
    77,
    82,
    83,
  }

  fmt.Println(average(xs))

}

/*
  Function to calculate averages.
  Argument: array of float64
  Returns: float64
 */
func average(xs []float64) float64 {
  // Variable to hold the total.
  total := 0.0

  // Iterate over the array of float64, and calculate the total.
  for _, v := range xs {
    total += v
  }

  // Return the average.
  return total / float64(len(xs))
}
