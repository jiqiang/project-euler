package main

import (
	"fmt"
	"github.com/jiqiang/project-euler/p10/lib"
)

func main() {
	sum := 2
	for i := 3; i <= 2000000; i += 2 {
		if lib.IsPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
