package main

import "fmt"

func main() {
  nextEven := makeEvenGenerator()

  fmt.Println(nextEven()) // 0
  fmt.Println(nextEven()) // 2
  fmt.Println(nextEven()) // 4
}

func makeEvenGenerator() func() uint {
  i := uint(0)

  return func() (ret uint) {
    ret = i
    i += 2
    return
  }
}
