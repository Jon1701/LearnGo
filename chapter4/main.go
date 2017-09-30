package main

import "fmt"

func main() {
  // Variable declaration and assignment.
  var x string = "Hello World"
  fmt.Println(x);

  // Variable reassignment.
  x = "Hello World"
  fmt.Println(x)

  // String concatenation.
  x += "I have been concatenated after Hello World"
  fmt.Println(x)

  // Variable declaration, and then assignment.
  var y string
  y = "Hello World"
  fmt.Println(y)

  // String equality.
  var s string = "hello";
  var t string = "world";
  fmt.Println(s == t)

  // Type inference (Go determined q is of type string).
  q := "Hello World"
  fmt.Println(q)

  // Type inference (Go determined r is of type int).
  r := 5
  fmt.Println(r);
}
