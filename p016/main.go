package main

import "fmt"
import "math"
import "strconv"

func main() {

	var s string

	sum := 0

	a := math.Pow(2, 1000)
	s = strconv.FormatFloat(a, 'f', 1, 64)

	for _, c := range s {
		cstr := string(c)
		if cstr == "." {
			break
		}
		v, _ := strconv.Atoi(cstr)
		sum += v
	}

	fmt.Println(sum)
}
