package main

import "fmt"

func main() {
  fmt.Println(add(1,2,3,4,5,6,7,8,9))
}

/*
  Variadic function.

  Argument `args` is of type `...int`, which means `args` is an int
  with zero or more values. (type variadic int, multiple int)
 */
func add(args ...int) int {
  // Variable to hold the total.
  total := 0

  // Calculate the total of all numbers in array `args`.
  for _, v := range args {
    total += v
  }

  // Return the total.
  return total
}
