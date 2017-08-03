//http://www.mozartreina.com/counting-lattice-paths.html
package main

import (
	"fmt"
)

func getNextRow(row []int) []int {
	var newRow = []int{}
	var n int
	for i := 1; i < len(row); i++ {
		if i == 1 {
			n = row[i] * 2
			newRow = append(newRow, n)
		} else {
			n = newRow[len(newRow)-1] + row[i]
			newRow = append(newRow, n)
		}
	}

	if len(newRow) > 1 {
		newRow = getNextRow(newRow)
	}

	return newRow
}

func main() {

	n := 21

	var start = []int{}

	for i := 0; i < n; i++ {
		start = append(start, 1)
	}

	row := getNextRow(start)
	fmt.Println(row[0])

}
