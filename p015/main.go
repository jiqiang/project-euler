package main

import "fmt"

type node struct {
	x int
	y int
}

func main() {

	var nodes, tempNodes []node
	n := node{0, 0}
	max := 3
	tempNodes = append(tempNodes, n)

	for i := 1; i <= max; i++ {
		for _, _n := range tempNodes {

			if _n.x < max {
				nn := node{_n.x + 1, _n.y}
				nodes = append(nodes, nn)
			}

			if _n.y < max {
				nn := node{_n.x, _n.y + 1}
				nodes = append(nodes, nn)
			}
		}
		tempNodes = nodes
		nodes = nil
	}

	fmt.Println(n)
	fmt.Println(nodes)
	fmt.Println(tempNodes)
	fmt.Println(max)
}
