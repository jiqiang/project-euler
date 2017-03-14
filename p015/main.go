package main

import (
	"fmt"
	"time"
)

type node struct {
	x int
	y int
}

const limit int = 10

func process(id int, in chan node) chan node {
	out := make(chan node, 1048576)
	go func() {
		for n := range in {

			fmt.Println(id, len(in))

			if n.x < limit {
				out <- node{n.x + 1, n.y}
			}

			if n.y < limit {
				out <- node{n.x, n.y + 1}
			}

		}
	}()
	return out
}

func main() {

	n := node{0, 0}
	in := make(chan node, 1048576)

	for i := 0; i < 20; i++ {
		out := process(1, in)
		in = out
		time.Sleep(1 * time.Second)
	}

	in <- n

	<-time.After(time.Second * 60)
}
