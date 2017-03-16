package main

import (
	"fmt"
	"runtime"
	"time"
)

type node struct {
	x int
	y int
}

const limit int = 20

func process(id int, n node) {
	tube := make(chan node, 904857600)

	go func() {
		for {

			n1 := <-tube
			//fmt.Printf("%d-1 %d\n", id, len(tube))
			if n1.x == limit || n1.y == limit {
				fmt.Println(id, len(tube))
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
			//fmt.Printf("%d-2 %d\n", id, len(tube))
			if n2.x == limit || n2.y == limit {
				fmt.Println(id, len(tube))
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

	//process(1, node{0, 0})
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	process(1, node{3, 1})
	process(2, node{2, 2})
	process(3, node{2, 1})
	process(4, node{1, 2})
	process(5, node{0, 3})

	fmt.Println()

	<-time.After(time.Minute * 3)

}
