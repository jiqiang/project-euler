package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type tree struct {
	left  *tree
	value int
	right *tree
}

func main() {
	file, err := os.Open("./data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
