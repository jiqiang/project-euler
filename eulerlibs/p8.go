package eulerlibs

import (
	"strconv"
	"strings"
)

// RunP8 runs problem 8 solution.
func RunP8(s string) int64 {

	numOfDigits := 13

	var strs []string

	var tempStr string

	for i := 0; i < len(s); i++ {

		tempStr = s[i : i+numOfDigits]

		if !strings.Contains(tempStr, "0") {
			strs = append(strs, tempStr)
		}

		if i+numOfDigits == len(s) {
			break
		}
	}

	var max, p int64 = 0, 1

	for _, str := range strs {

		p = 1

		for _, r := range str {
			d, _ := strconv.ParseInt(string(r), 10, 64)

			p *= d
		}

		if max < p {
			max = p
		}
	}

	return max
}
