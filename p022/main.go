package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

// LetterPositionMap represents a map which key is letter and value is position.
var LetterPositionMap = map[string]int{
	"A": 1, "B": 2, "C": 3, "D": 4,
	"E": 5, "F": 6, "G": 7, "H": 8,
	"I": 9, "J": 10, "K": 11, "L": 12,
	"M": 13, "N": 14, "O": 15, "P": 16,
	"Q": 17, "R": 18, "S": 19, "T": 20,
	"U": 21, "V": 22, "W": 23, "X": 24,
	"Y": 25, "Z": 26,
}

// ExtractNamesFromString extracts names from giving string into a slice with
// names sorted.
func ExtractNamesFromString(str string) []string {
	names := []string{}
	cleanStr := strings.Trim(str, " ")
	if len(cleanStr) == 0 {
		return names
	}
	for _, name := range strings.Split(cleanStr, ",") {
		cleanName := strings.Trim(name, " ")
		names = append(names, cleanName[1:len(cleanName)-1])
	}
	sort.Strings(names)
	return names
}

// GetContentStringFromFile gets content from a file without leading and ending
// spaces
func GetContentStringFromFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(content)), nil
}

// GetLetterPositionSum sums all letters alphabetical position.
func GetLetterPositionSum(name string) int64 {
	sum := int64(0)
	for _, l := range name {
		p, ok := LetterPositionMap[string(l)]
		if ok {
			sum += int64(p)
		}
	}
	return sum
}

func main() {
	content, err := GetContentStringFromFile("names.txt")
	if err != nil {
		log.Fatal(err)
	}
	names := ExtractNamesFromString(content)
	sum := int64(0)
	for index, name := range names {
		posSum := GetLetterPositionSum(name)
		sum += int64(posSum) * int64(index+1)
	}
	fmt.Println(sum)
}
