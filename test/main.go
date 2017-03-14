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

func process(id int, tube chan node, count *int) {

	for {
		n := <-tube

		fmt.Println(id, len(tube))

		if n.x < limit {
			tube <- node{n.x + 1, n.y}
		}

		if n.y < limit {
			tube <- node{n.x, n.y + 1}
		}

		if n.x == limit && n.y == limit {
			*count++
		}
		time.Sleep(1 * time.Second)
	}

}

func main() {

	n := node{0, 0}
	tube := make(chan node, 100)
	count1, count2, count3, count4, count5 := 0, 0, 0, 0, 0

	go process(1, tube, &count1)
	go process(2, tube, &count2)
	go process(3, tube, &count3)
	go process(4, tube, &count4)
	go process(5, tube, &count5)

	tube <- n

	<-time.After(time.Second * 60)
	fmt.Println(count1 + count2 + count3 + count4 + count5)
}
