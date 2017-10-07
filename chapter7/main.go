package main

import "fmt"

func main() {
  defer second()
  first()

  defer func() {
    str := recover()
    fmt.Println(str)
  }()

  panic("PANIC")
}

func first() {
  fmt.Println("1st")
}

func second() {
  fmt.Println("2nd")
}
