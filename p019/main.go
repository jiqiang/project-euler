package main

import "fmt"
import "time"

func calculateInYear(year int, sundays chan int) {
	sum := 0
	for m := 1; m <= 12; m++ {
		t := time.Date(year, time.Month(m), 1, 0, 0, 0, 0, time.UTC)
		if t.Weekday().String() == "Sunday" {
			sum++
		}
	}
	sundays <- sum
}

func main() {

	numOfSundays := 0
	sundaysChan := make(chan int)

	for y := 1901; y <= 2000; y++ {
		go calculateInYear(y, sundaysChan)
	}

	for y := 1901; y <= 2000; y++ {
		numOfSundays = numOfSundays + <-sundaysChan
	}

	fmt.Println(numOfSundays)
}
