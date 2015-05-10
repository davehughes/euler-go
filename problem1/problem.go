package problem1

func divisibleByAll(i int, divs []int) bool {
	for _, item := range divs {
		if i%item != 0 {
			return false
		}
	}
	return true
}

func divisibleByAny(i int, divs []int) bool {
	for _, item := range divs {
		if i%item == 0 {
			return true
		}
	}
	return false
}

func genrange(start int, end int) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < end; i++ {
			// fmt.Printf("Generating %d\n", i)
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func filter(in chan int, out chan int, f func(int) bool) {
	for item := range in {
		if f(item) {
			out <- item
		}
	}
	close(out)
}

func Run() int {
	divs := []int{3, 5}

	range_chan := genrange(0, 1000)
	divisible_chan := make(chan int)
	go filter(range_chan, divisible_chan, func(i int) bool {
		return divisibleByAny(i, divs)
	})

	sum := 0
	for x := range divisible_chan {
		sum += x
	}
	return sum
}
