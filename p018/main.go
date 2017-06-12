package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type tree struct {
	left  *tree
	value int
	right *tree
}

func stringSliceToIntSlice(sString []string) []int {
	var sInt = []int{}
	for _, s := range sString {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		sInt = append(sInt, i)
	}
	return sInt
}

func getMatrixFromDataFile(filePath string) ([][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matrix = [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r := strings.Fields(scanner.Text())
		intSlice := stringSliceToIntSlice(r)
		matrix = append(matrix, intSlice)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return matrix, nil
}

func displayMatrix(matrix [][]int) {
	for _, r := range matrix {
		fmt.Println(r)
	}
}

func main() {
	matrix, err := getMatrixFromDataFile("./data")
	if err != nil {
		log.Fatal(err)
	}

	//displayMatrix(matrix)

	sum, index := 0, 0

	for i, r := range matrix {

		if i == 0 {
			sum += r[index]
			continue
		}

		if r[index] >= r[index+1] {
			sum += r[index]
			fmt.Println(r[index])
		} else {
			sum += r[index+1]
			fmt.Println(r[index+1])
			index++
		}
	}

	fmt.Println(sum)

}
