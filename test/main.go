package main

import (
	"fmt"
	"time"
)

type node struct {
	x int
	y int
}

func process(tube chan node) {
	for {
		n := <-tube
		fmt.Println(n)
		time.Sleep(100 * time.Millisecond)
		if n.x == -1 && n.y == -1 {
			tube <- node{-1, -1}
		} else if n.x < 2 {
			tube <- node{n.x + 1, n.y}
		} else if n.y < 2 {
			tube <- node{n.x, n.y + 1}
		} else {
			tube <- node{-1, -1}
		}
	}
}

func main() {

	n := node{0, 0}
	tube := make(chan node)

	go process(tube)
	go process(tube)

	tube <- n
	time.Sleep(1 * time.Second)
	<-tube
}
