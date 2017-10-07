package main

import "fmt"

func main() {
  x, y, z := f(2, 3, 1)

  fmt.Println(x, y, z)
}

func f(x int, y int, z int) (int, int, int) {
  return x, y, z
}
