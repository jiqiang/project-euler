package main

import (
	"fmt"
)

type node struct {
	x int
	y int
}

func process(n node) uint64 {
	if n.x == 0 && n.y == 0 {
		return 1
	}

	var paths uint64

	paths = 0

	if n.x > 0 {
		paths += process(node{n.x - 1, n.y})
	}

	if n.y > 0 {
		paths += process(node{n.x, n.y - 1})
	}

	return paths
}

func main() {

	result := process(node{2, 2})

	fmt.Println(result)

}
