package problem5

func divisibleByAll(i int, divs []int) bool {
	for _, item := range divs {
		if i%item != 0 {
			return false
		}
	}
	return true
}

func Run() int {
	ns := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	pred := func(i int) bool {
		return divisibleByAll(i, ns)
	}
	for i := 19; ; i += 19 {
		if pred(i) {
			return i
		}
	}
}
