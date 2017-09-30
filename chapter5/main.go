package main

import "fmt"

func main() {
  // Set `i` to 1. Infer type int.
  i := 1

  // Loop from 1 to 10, inclusive.
  for i <= 10 {
    fmt.Println(i)

    // Increment `i`
    i = i + 1
  }

  // Another way to write a for-loop.
  for i := 1; i <= 10; i++ {
    fmt.Println(i)
  }
}
