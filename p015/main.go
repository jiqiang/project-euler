package main

import "fmt"

type node struct {
	x int
	y int
}

func done(ns []node, max int) (bool, int) {
	count := 0
	for _, n := range ns {
		if n.x < max || n.y < max {
			return false, count
		}
		count++
	}
	return true, count
}

func main() {

	max := 10

	n := node{0, 0}

	var ns1, ns2 []node

	ns1 = append(ns1, n)

	for {
		for _, _n := range ns1 {
			if _n.x < max {
				tn := node{_n.x + 1, _n.y}
				ns2 = append(ns2, tn)
			}

			if _n.y < max {
				tn := node{_n.x, _n.y + 1}
				ns2 = append(ns2, tn)
			}
		}

		done, count := done(ns2, max)

		fmt.Println(len(ns2))
		//fmt.Println(ns2)

		if done {
			fmt.Println(count)
			break
		}

		ns1 = ns2
		ns2 = nil
	}

}
