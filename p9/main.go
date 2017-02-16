package main

import (
  "fmt"
  "github.com/jiqiang/project-euler/p9/lib"
)

func main() {
  x, y, z, product := lib.P9()
  fmt.Printf("%d * %d * %d = %d\n", x, y, z, product)
}
