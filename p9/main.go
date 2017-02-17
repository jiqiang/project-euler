package main

import (
	"fmt"
	"github.com/jiqiang/project-euler/eulerlibs"
)

func main() {
	x, y, z, product := eulerlibs.P9()
	fmt.Printf("%d * %d * %d = %d\n", x, y, z, product)
}
