package main

import (
	"fmt"
	"time"
)

type node struct {
	x int
	y int
}

const limit int = 3

func process(id int, in chan node) chan node {
	out := make(chan node, 1048576)
	go func() {
		for n := range in {

			if n.x < limit {
				out <- node{n.x + 1, n.y}
			}

			if n.y < limit {
				out <- node{n.x, n.y + 1}
			}

			//time.Sleep(1 * time.Second)
		}
		close(out)
	}()
	return out
}

func main() {

	n := node{0, 0}
	in := make(chan node, 1048576)
	count := 0

	out1 := process(1, in)
	out2 := process(2, out1)
	out3 := process(2, out2)
	out4 := process(2, out3)

	in <- n

	go func() {
		for {
			_n := <-out4
			fmt.Println(_n)
			if _n.x == limit && _n.y == limit {
				count++
			}
		}
	}()

	<-time.After(time.Second * 10)

}
