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

func process(id int, n node) {
	tube := make(chan node, 1048576)
	go func() {
		for {
			n1 := <-tube

			if n1.x == limit || n1.y == limit {
				fmt.Println(id, n1)
				continue
			}

			if n1.x < limit {
				tube <- node{n1.x + 1, n1.y}
			}

			if n1.y < limit {
				tube <- node{n1.x, n1.y + 1}
			}
		}
	}()

	go func() {
		for {
			n2 := <-tube

			if n2.x == limit || n2.y == limit {
				fmt.Println(id, n2)
				continue
			}

			if n2.x < limit {
				tube <- node{n2.x + 1, n2.y}
			}

			if n2.y < limit {
				tube <- node{n2.x, n2.y + 1}
			}
		}
	}()

	tube <- n
}

func main() {

	process(1, node{2, 0})
	process(2, node{1, 1})
	process(3, node{1, 1})
	process(4, node{0, 2})

	<-time.After(time.Second * 10)

}
