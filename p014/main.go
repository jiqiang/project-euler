package main

import "fmt"

func main() {
	var seqs [][]int
	m := make(map[int]int)

	for n := 2; n <= 13; n++ {
		var seq []int
		sn := n
		count := 1
		seq = append(seq, n)
		for {

			if sn == 1 {
				break
			}

			if sn%2 == 0 {
				sn = sn / 2
			} else {
				sn = sn*3 + 1
			}

			seq = append(seq, sn)
			count++
		}
		seqs = append(seqs, seq)
		m[n] = count
	}
	fmt.Println(seqs)
	fmt.Println(m)
}
