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

func getRowWithMax(row []int) []int {
	var newRow = []int{}
	for i := 0; i < len(row); i++ {
		if i+1 < len(row) {
			if row[i] >= row[i+1] {
				newRow = append(newRow, row[i])
			} else {
				newRow = append(newRow, row[i+1])
			}
		}
	}
	return newRow
}

func main() {
	matrix, err := getMatrixFromDataFile("./data")
	if err != nil {
		log.Fatal(err)
	}

	//displayMatrix(matrix)

	var rowWithMax, lastRowWithMax []int

	lastRowWithMax = getRowWithMax(matrix[len(matrix)-1])

	numOfRows := len(matrix)

	for i := numOfRows - 2; i > 0; i-- {
		rowWithMax = nil
		for j := 0; j < len(matrix[i]); j++ {
			rowWithMax = append(rowWithMax, lastRowWithMax[j]+matrix[i][j])
		}

		lastRowWithMax = getRowWithMax(rowWithMax)
	}

	fmt.Println(lastRowWithMax[0] + matrix[0][0])

}
