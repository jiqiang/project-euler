/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
package main

import "fmt"
import "sync"
import "sort"

const TARGET = 600851475143
const COUNT = 10000

func isPrime(num int) bool {
  if num <= 1 {
    return false
  }

  if num == 2 {
    return true
  }

  for i := 2; i < num; i++ {
    if num % i == 0 {
      return false
    }
  }

  return true
}

func getFactor(target int, primes []int, factors *[]int) {
  for _, v := range primes {
    if target % v == 0 {

      if target / v == 1 {
        *factors = append(*factors, v)
        return
      }

      *factors = append(*factors, v)
      getFactor(target / v, primes, factors)
      return
    }
  }
}

func worker(start int, end int, results chan<- int, wg *sync.WaitGroup) {
  for i := start; i < end; i++ {
    if isPrime(i) {
      results <- i
    }
  }
  wg.Done()
}

func main() {

  var primes, factors []int
  var wg sync.WaitGroup

  results := make(chan int, 10000)

  for i := 1; i < 10; i++ {
    wg.Add(1)
    start := i * COUNT - COUNT
    end := i * COUNT
    if end > TARGET {
      end = TARGET
    }
    go worker(start, end, results, &wg)
  }

  wg.Wait()

  for i := 0; i < len(results); i++ {
    primes = append(primes, <-results)
  }

  sort.Ints(primes)

  getFactor(TARGET, primes, &factors)

  fmt.Println(factors)

}
