package main

import "fmt"

func main() {
	m := make(map[int]int)
	mc := 0
	var lsn int
	for n := 2; n <= 1000000; n++ {
		var seq []int
		sn := n
		seq = append(seq, n)
		count := 1
		for {
			if sn == 1 {
				break
			}
			if sn%2 == 0 {
				sn = sn / 2
			} else {
				sn = sn*3 + 1
			}
			if c, ok := m[sn]; ok {
				count += c
				break
			}
			seq = append(seq, sn)
			count++
		}
		if mc < count {
			mc = count
			lsn = n
		}
		for idx, v := range seq {
			_, ok := m[v]
			if !ok {
				m[v] = count - idx
			}
		}
	}
	fmt.Println(lsn)
}
