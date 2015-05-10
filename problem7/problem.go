package problem7

import "github.com/davehughes/euler-go/util"

func Run() int {
	primes_chan := util.GenPrimes()
	idx := 0
	for prime := range primes_chan {
		idx++
		if idx == 10001 {
			close(primes_chan)
			return prime
		}
	}
	return -1
}
