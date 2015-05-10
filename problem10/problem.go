package problem10

import "github.com/davehughes/euler-go/util"

func Run() int {
	sum := 0
	upperLimit := 2000000
	primesChan := util.GenPrimes()
	for prime := range primesChan {
		if prime > upperLimit {
			close(primesChan)
			break
		}
		sum += prime
	}
	return sum
}
