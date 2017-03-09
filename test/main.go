package main

import (
	"fmt"
)

type node struct {
	x int
	y int
}

func display(n node, c chan node) {
	go func() {
		c <- n
	}()
}

func main() {

	n := node{0, 0}

	c := make(chan node)

	display(n, c)

	fmt.Println(<-c)

}
