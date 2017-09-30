package main

import "fmt"

func main() {
  slice1 := []int{1,2,3}

  /*
  `append` creates a new slice by taking the existing slice `slice1`
   and appending all the following arguments to it.
  */
  slice2 := append(slice1, 4, 5)
  fmt.Println(slice1, slice2)

  // Create a slice of size 2.
  slice3 := make([]int, 2)

  // Copy elements from `slice1` into `slice3`
  copy(slice3, slice1)

  // Only the first 2 elements are copied, since `slice3` has a size of 2.
  fmt.Println(slice3);
}
