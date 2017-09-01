package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func getNameList(file string) ([]string, error) {
	data, err := ioutil.ReadFile("names.txt")
	if err != nil {
		return nil, err
	}
	names := []string{}
	for _, line := range strings.Split(string(data), ",") {
		names = append(names, line[1:len(line)-1])
	}
	sort.Strings(names)
	return names, nil
}

func getLetterPositonSum(name string, lpm map[string]int) int64 {
	sum := int64(0)
	for _, l := range name {
		p, ok := lpm[string(l)]
		if ok {
			sum += int64(p)
		}
	}
	return sum
}

func main() {
	letterPositionMap := map[string]int{
		"A": 1, "B": 2, "C": 3, "D": 4,
		"E": 5, "F": 6, "G": 7, "H": 8,
		"I": 9, "J": 10, "K": 11, "L": 12,
		"M": 13, "N": 14, "O": 15, "P": 16,
		"Q": 17, "R": 18, "S": 19, "T": 20,
		"U": 21, "V": 22, "W": 23, "X": 24,
		"Y": 25, "Z": 26,
	}

	names, err := getNameList("names.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := int64(0)
	for idx, name := range names {
		sum := getLetterPositonSum(name, letterPositionMap)
		result += sum * int64(idx+1)
	}
	fmt.Println(result)
}
