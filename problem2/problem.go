package problem2

import "github.com/davehughes/euler-go/util"

func IsEven(n int) bool {
	return n%2 == 0
}

func Run() int {
	sum := 0
	even_fib_chan := util.MakeFilteredChan(util.GenFib(), IsEven)

	for item := range even_fib_chan {
		if item >= 4000000 {
			break
		}
		sum += item
	}
	close(even_fib_chan)
	return sum
}
