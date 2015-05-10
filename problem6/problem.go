package problem6

func SumOfSquares(ns []int) int {
	s := 0
	for _, n := range ns {
		s += n * n
	}
	return s
}

func SquareOfSum(ns []int) int {
	s := 0
	for _, n := range ns {
		s += n
	}
	return s * s
}

func Run() int {
	numbers := 100
	ns := make([]int, numbers)
	for i := 0; i < numbers; i++ {
		ns[i] = i + 1
	}

	return SquareOfSum(ns) - SumOfSquares(ns)
}
