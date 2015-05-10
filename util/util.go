package util

import "math"

func Max(ns []int) int {
	max := math.MinInt64
	for _, n := range ns {
		if n > max {
			max = n
		}
	}
	return int(max)
}

func Sum(ch chan int) int {
	sum := 0
	for x := range ch {
		sum += x
	}
	return sum
}

func GenRange(start int, end int) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < end; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func MakeFilteredChan(ch chan int, f func(int) bool) chan int {
	out := make(chan int)
	go func() {
		for item := range ch {
			if f(item) {
				out <- item
			}
		}
		close(out)
	}()
	return out
}

func Take(n int, ch chan int) []int {
	items := make([]int, n)
	idx := 0

	for item := range ch {
		items[idx] = item

		idx++
		if idx >= n {
			break
		}
	}
	return items
}

func GenFib() chan int {
	ch := make(chan int, 1)
	f1 := 0
	f2 := 1
	go func() {
		for {
			f1, f2 = f2, f1+f2
			ch <- f1
		}
	}()
	return ch
}

func GenPrimes() chan int {
	out := make(chan int)

	primesSoFar := []int{}
	i := 2

	go func() {
		for {
			isPrime := true
			for _, p := range primesSoFar {
				if i%p == 0 {
					isPrime = false
					break
				}
			}

			if isPrime {
				primesSoFar = append(primesSoFar, i)
				// fmt.Println("Found prime:", i)
				out <- i
			}
			i += 1
		}
	}()

	return out
}
