package main

import (
	"fmt"
	"time"
)

type node struct {
	x int
	y int
}

func process(in chan *node, out chan *node) {
	go func() {
		for {
			n := <-in
			fmt.Println(*n)
			if n.x < 2 {
				out <- &node{n.x + 1, n.y}
			}

			if n.y < 2 {
				out <- &node{n.x, n.y + 1}
			}
		}
	}()
}

func main() {

	in, out := make(chan *node), make(chan *node)

	process(in, out)
	process(out, in)

	in <- &node{0, 0}

	time.Sleep(3 * time.Second)

}
