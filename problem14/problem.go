package problem14

import "fmt"

var memoizeCache = make(map[int]int, 1000000)

func CollatzLength(n int) int {
	memoized, ok := memoizeCache[n]
	if ok {
		return memoized
	}

	var length int
	if n == 1 {
		length = 1
	} else if n%2 == 0 {
		length = 1 + CollatzLength(n/2)
	} else {
		length = 1 + CollatzLength(3*n+1)
	}
	memoizeCache[n] = length
	return length
}

func Run() int {
	maxLen := 1
	maxNum := 1
	limit := 1000000
	for i := 1; i < limit; i++ {
		fmt.Println("Calcing collatz:", i)
		clen := CollatzLength(i)
		if clen > maxLen {
			maxNum = i
			maxLen = clen
		}
	}
	return maxNum
}
