package main

import "fmt"

func main() {
  // Declare and assign a constant.
  const x string = "Hello World"
  fmt.Println(x)

  // Attempt to re-assign x.
  x = "Some other string"

  /*
    Should display the following error:
    ./main.go:11:5: cannot assign to x
  */
}
