package main

import "fmt"

/*
 Variable `x` is outside of `main`.
 Other functions are able to access this variable.
*/
var x string = "Hello World"

func main() {
  fmt.Println(x)

  f()
}

// This function has access to `x`
func f() {
  fmt.Println(x)
}
