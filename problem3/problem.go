package problem3

import "github.com/davehughes/euler-go/util"

// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

func FindPrimeFactors(n int) []int {
	factors := []int{}

	for {
		primesChan := util.GenPrimes()
		for p := range primesChan {
			if n%p == 0 {
				factors = append(factors, p)
				n = n / p
				break
			}
		}

		if n == 1 {
			break
		}
	}

	return factors
}

func Run() int {
	return util.Max(FindPrimeFactors(600851475143))
}
