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

func process(id int, tube chan node) {
	for {
		n := <-tube

		fmt.Println(id, n)
		//fmt.Println(len(tube))

		if n.x < limit {
			tube <- node{n.x + 1, n.y}
		}

		if n.y < limit {
			tube <- node{n.x, n.y + 1}
		}
	}
}

func main() {

	tube := make(chan node, 100)

	go process(1, tube)
	go process(2, tube)
	go process(3, tube)

	tube <- node{0, 0}

	<-time.After(time.Second * 15)

}
