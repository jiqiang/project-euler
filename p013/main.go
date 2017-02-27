package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./p013/p013")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var sum uint64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.ParseUint(scanner.Text()[0:11], 10, 64)
		if err != nil {
			panic(err)
		}
		sum += num
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
