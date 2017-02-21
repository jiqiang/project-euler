package main

import "fmt"
import "github.com/jiqiang/project-euler/eulerlibs"

func main() {
	sum := 0
	for i := 1; ; i++ {
		sum = sum + i
		if eulerlibs.IsPrime2(sum) {
			continue
		}
		count := 0
		for j := 1; j <= sum; j++ {
			if sum%j == 0 {
				count++
			}
		}
		fmt.Println(count)
		if count > 100 {
			break
		}

	}
	fmt.Println(sum)
}
