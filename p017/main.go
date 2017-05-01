/*
If the numbers 1 to 5 are written out in words: one, two, three, four, five,
then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in
words, how many letters would be used?

NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and
forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20
letters. The use of "and" when writing out numbers is in compliance with
British usage.
*/

package main

import "fmt"
import "strconv"

func letterCount(n int, m map[string]int, mm map[string]int) int {
	count := 0
	s := strconv.Itoa(n)

	if n < 20 {
		count += m[s]
	} else if n < 100 {
		if n%10 == 0 {
			count += mm[string(s[0])]
		} else {
			count += mm[string(s[0])] + m[string(s[1])]
		}
	} else if n < 1000 {
		if n%100 == 0 {
			count += m[string(s[0])] + m["hundred"]
		} else if string(s[1]) == "0" {
			count += m[string(s[0])] + m["hundred"] + m["and"] + m[string(s[2])]
		} else if string(s[1]) == "1" {
			count += m[string(s[0])] + m["hundred"] + m["and"] + m[string(s[1:])]
		} else {
			count += m[string(s[0])] + m["hundred"] + m["and"] + mm[string(s[1])] + m[string(s[2])]
		}
	} else {
		count += m[string(s[0])] + m["thousand"]
	}

	return count
}

func main() {
	m := make(map[string]int)
	mm := make(map[string]int)

	m["1"] = 3  // one
	m["2"] = 3  // two
	m["3"] = 5  // three
	m["4"] = 4  // four
	m["5"] = 4  // five
	m["6"] = 3  // six
	m["7"] = 5  // seven
	m["8"] = 5  // eight
	m["9"] = 4  // nine
	m["10"] = 3 // ten
	m["11"] = 6 // eleven
	m["12"] = 6 // twelve
	m["13"] = 8 // thirteen
	m["14"] = 8 // fourteen
	m["15"] = 7 // fifteen
	m["16"] = 7 // sixteen
	m["17"] = 9 // seventeen
	m["18"] = 8 // eighteen
	m["19"] = 8 // nineteen
	mm["2"] = 6 // twenty
	mm["3"] = 6 // thirty
	mm["4"] = 5 // forty
	mm["5"] = 5 // fifty
	mm["6"] = 5 // sixty
	mm["7"] = 7 // seventy
	mm["8"] = 6 // eighty
	mm["9"] = 6 // ninety
	m["hundred"] = 7
	m["thousand"] = 8
	m["and"] = 3

	//count := letterCount(342, m, mm)
	count := 0
	for n := 1; n <= 1000; n++ {
		count += letterCount(n, m, mm)
	}

	fmt.Println(count)
}
