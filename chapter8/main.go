package main

import "fmt"

func zero(xPtr *int) {
  // * dereference: give access to the value which xPtr points to
  *xPtr = 0
}

func main() {
  x := 5

  // & returns a pointer (to an int)
  zero(&x)

  fmt.Println(x) // `x` is 0
}
