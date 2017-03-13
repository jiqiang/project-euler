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
		tube <- node{n.x + 1, n.y}

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
