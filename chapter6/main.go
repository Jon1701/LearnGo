package main

import "fmt"

func main() {
  /*
  Create a slice which is associated with an underlying
  float64 array of length 5.
  */
  x := make([]float64, 5, 10)
  fmt.Println(x)

  /*
  Another way to create slices is to use the [low:high]
  expression.
  */
  arr := []float64{1,2,3,4,5}

  /*
  Create a slice using [low:high expressions].
  Can omit endpoints.
  */
  y := arr[0:5]
  z := arr[0:]
  u := arr[:5]
  v := arr[:]

  fmt.Println(y)
  fmt.Println(z)
  fmt.Println(u)
  fmt.Println(v)
}
