package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Tree represents a binary tree structure.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
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
	matrix, err := getMatrixFromDataFile("./smalldata")
	if err != nil {
		log.Fatal(err)
	}

	displayMatrix(matrix)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i+1 == len(matrix) {
				fmt.Println(matrix[i][j])
			} else {
				fmt.Println(matrix[i][j], matrix[i+1][j], matrix[i+1][j+1])
			}
		}
	}
}
