package problem11

var grid = [][]int{
	[]int{8, 2, 22, 97, 38, 15, 0, 40, 0, 75, 4, 5, 7, 78, 52, 12, 50, 77, 91, 8},
	[]int{49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 4, 56, 62, 0},
	[]int{81, 49, 31, 73, 55, 79, 14, 29, 93, 71, 40, 67, 53, 88, 30, 3, 49, 13, 36, 65},
	[]int{52, 70, 95, 23, 4, 60, 11, 42, 69, 24, 68, 56, 1, 32, 56, 71, 37, 02, 36, 91},
	[]int{22, 31, 16, 71, 51, 67, 63, 89, 41, 92, 36, 54, 22, 40, 40, 28, 66, 33, 13, 80},
	[]int{24, 47, 32, 60, 99, 3, 45, 2, 44, 75, 33, 53, 78, 36, 84, 20, 35, 17, 12, 50},
	[]int{32, 98, 81, 28, 64, 23, 67, 10, 26, 38, 40, 67, 59, 54, 70, 66, 18, 38, 64, 70},
	[]int{67, 26, 20, 68, 2, 62, 12, 20, 95, 63, 94, 39, 63, 8, 40, 91, 66, 49, 94, 21},
	[]int{24, 55, 58, 05, 66, 73, 99, 26, 97, 17, 78, 78, 96, 83, 14, 88, 34, 89, 63, 72},
	[]int{21, 36, 23, 9, 75, 0, 76, 44, 20, 45, 35, 14, 0, 61, 33, 97, 34, 31, 33, 95},
	[]int{78, 17, 53, 28, 22, 75, 31, 67, 15, 94, 03, 80, 4, 62, 16, 14, 9, 53, 56, 92},
	[]int{16, 39, 5, 42, 96, 35, 31, 47, 55, 58, 88, 24, 0, 17, 54, 24, 36, 29, 85, 57},
	[]int{86, 56, 0, 48, 35, 71, 89, 7, 5, 44, 44, 37, 44, 60, 21, 58, 51, 54, 17, 58},
	[]int{19, 80, 81, 68, 5, 94, 47, 69, 28, 73, 92, 13, 86, 52, 17, 77, 4, 89, 55, 40},
	[]int{4, 52, 8, 83, 97, 35, 99, 16, 07, 97, 57, 32, 16, 26, 26, 79, 33, 27, 98, 66},
	[]int{88, 36, 68, 87, 57, 62, 20, 72, 3, 46, 33, 67, 46, 55, 12, 32, 63, 93, 53, 69},
	[]int{4, 42, 16, 73, 38, 25, 39, 11, 24, 94, 72, 18, 8, 46, 29, 32, 40, 62, 76, 36},
	[]int{20, 69, 36, 41, 72, 30, 23, 88, 34, 62, 99, 69, 82, 67, 59, 85, 74, 4, 36, 16},
	[]int{20, 73, 35, 29, 78, 31, 90, 01, 74, 31, 49, 71, 48, 86, 81, 16, 23, 57, 5, 54},
	[]int{1, 70, 54, 71, 83, 51, 54, 69, 16, 92, 33, 48, 61, 43, 52, 1, 89, 19, 67, 48},
}

func HasHorizontal(x, y, w, h, size int) bool {
	return x+size < w
}

func HasVertical(x, y, w, h, size int) bool {
	return y+size < h
}

func HasDiagRight(x, y, w, h, size int) bool {
	return (x+size < w) && (y+size < h)
}

func HasDiagLeft(x, y, w, h, size int) bool {
	return (x-size >= 0) && (y+size < h)
}

func GetHorizontal(x, y, w, h, size int) []int {
	if !HasHorizontal(x, y, w, h, size) {
		return []int{}
	}

	group := make([]int, size)
	for i := 0; i < size; i++ {
		group[i] = grid[y][x+i]
	}
	return group
}

func GetVertical(x, y, w, h, size int) []int {
	if !HasVertical(x, y, w, h, size) {
		return []int{}
	}

	group := make([]int, size)
	for i := 0; i < size; i++ {
		group[i] = grid[y+i][x]
	}
	return group
}

func GetDiagRight(x, y, w, h, size int) []int {
	if !HasDiagRight(x, y, w, h, size) {
		return []int{}
	}

	group := make([]int, size)
	for i := 0; i < size; i++ {
		group[i] = grid[y+i][x+i]
	}
	return group
}

func GetDiagLeft(x, y, w, h, size int) []int {
	if !HasDiagLeft(x, y, w, h, size) {
		return []int{}
	}

	group := make([]int, size)
	for i := 0; i < size; i++ {
		group[i] = grid[y+i][x-i]
	}
	return group
}

func Product(ns []int) int {
	s := 1
	for _, n := range ns {
		s = s * n
	}
	return s
}

func Run() int {
	w := len(grid[0])
	h := len(grid)
	size := 4

	groupFuncs := []func(x, y, w, h, size int) []int{
		GetHorizontal, GetVertical, GetDiagRight, GetDiagLeft,
	}

	maxProduct := -1
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			for _, f := range groupFuncs {
				group := f(i, j, w, h, size)
				if len(group) == 0 {
					continue
				}

				product := Product(group)
				if product > maxProduct {
					maxProduct = product
				}
			}
		}
	}
	return maxProduct
}
