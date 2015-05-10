package problem12

import (
	"fmt"
	"math"
)

func GenTriangleNumbers() chan int {
	out := make(chan int)

	go func() {
		acc := 1
		idx := 1
		for {
			out <- acc
			idx++
			acc += idx
		}
	}()

	return out
}

func FindFactors(n int) []int {
	if n == 1 {
		return []int{1}
	}

	factors := []int{}
	for i := 1; float64(i) < math.Sqrt(float64(n)); i++ {
		if n%i == 0 {
			factors = append(factors, i, n/i)
		}
	}
	return factors
}

func Run() int {
	triangleNumbers := GenTriangleNumbers()

	for num := range triangleNumbers {
		factors := FindFactors(num)
		fmt.Println(num, "has", len(factors), "factors")
		if len(factors) > 500 {
			close(triangleNumbers)
			return num
		}
	}
	return -1
}
