package main

import "fmt"

func main() {
  // Bad name for a variable.
  x := "Max"
  fmt.Println("My dog's name is", x)

  // Better name for a variable.
  name := "Max"
  fmt.Println("My dog's name is", name)

  // Better name for a variable.
  dogsName := "Max"
  fmt.Println("My dog's name is", dogsName)
}
