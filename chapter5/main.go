package main

import "fmt"

func main() {
  i := 0

  // Print english names for numbers, using if statements.
  if i == 0 {
    fmt.Println("Zero")
  } else if i == 1 {
    fmt.Println("One")
  } else if i == 2 {
    fmt.Println("Two")
  } else if i == 3 {
    fmt.Println("Three")
  } else if i == 4 {
    fmt.Println("Four")
  } else if i == 5 {
    fmt.Println("Five")
  }

  i = 100

  // Print english names for numbers, using switch statement.
  switch i {
    case 0: fmt.Println("Zero")
    case 1: fmt.Println("One")
    case 2: fmt.Println("Two")
    case 3: fmt.Println("Three")
    case 4: fmt.Println("Four")
    case 5: fmt.Println("Five")
    default: fmt.Println("Unknown number")
  }
}
