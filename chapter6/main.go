package main

import "fmt"

func main() {
  // `x` is a mapping from strings to ints
  // `make` is used to initialize.
  x := make(map[string]int)

  // Assign 10 to the key 'key'.
  x["key"] = 10

  // Display value of 'key' in map `x`
  fmt.Println(x["key"]);

  // `y` is a mapping from ints to ints.
  y := make(map[int]int)

  // Assign 10 to y with key 1.
  y[1] = 10

  // Display value of 1 in map `y`
  fmt.Println(y[1])

  // Create another map from strings to float32.
  z := make(map[string]float32)

  // Assign values to map `z`
  z["pi"] = 3.141592654
  z["e"] = 2.7131

  // Delete `e`
  delete(z, "e")

  // Display map of `z`
  fmt.Println(z)

  // Add some more values to `z`.
  z["pythagoras"] = 1.4142135623
  z["theodorus"] = 1.7320508075
  z["golden"] = 1.6180339887
  z["omega"] = 0.567143290409783

  // Extract value of z at key 'omega', and check if it exists
  name, ok := z["omega"]

  fmt.Println(name, ok)

  // Shorter way to create maps.
  u := map[string]string{
    "H": "Hydrogen",
    "He": "Helium",
    "Li": "Lithium",
  }

  fmt.Println(u)

  // Nested map.
  v := map[string]map[string]string{
    "H": map[string]string{
      "name": "Hydrogen",
      "state": "gas",
    },
    "He": map[string]string{
      "name": "Helium",
      "state": "gas",
    },
    "Li": map[string]string{
      "name": "Lithium",
      "state": "solid",
    },
  }

  if el, ok := v["Li"]; ok {
    fmt.Println(el["name"], el["state"])
  }
}
