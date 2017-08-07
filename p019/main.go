package main

import "fmt"
import "time"

func main() {

	numOfSundays := 0

	for y := 1901; y <= 2000; y++ {
		for m := 1; m <= 12; m++ {
			t := time.Date(y, time.Month(m), 1, 1, 0, 0, 0, time.UTC)
			if t.Weekday().String() == "Sunday" {
				numOfSundays++
			}
		}
	}

	fmt.Println(numOfSundays)
}
