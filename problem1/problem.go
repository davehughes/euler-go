package problem1

import "github.com/davehughes/euler-go/util"

func divisibleByAny(i int, divs []int) bool {
	for _, item := range divs {
		if i%item == 0 {
			return true
		}
	}
	return false
}

func Run() int {
	divs := []int{3, 5}
	range_chan := util.GenRange(0, 1000)
	divisible_chan := util.MakeFilteredChan(range_chan, func(i int) bool {
		return divisibleByAny(i, divs)
	})
	return util.Sum(divisible_chan)
}
