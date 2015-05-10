package problem15

type Vertex struct {
	w, h int
}

var memoizeCache = make(map[Vertex]int)

func CountLatticePaths(w, h int) int {
	memoized, ok := memoizeCache[Vertex{w, h}]
	if ok {
		return memoized
	}

	if w == 0 || h == 0 {
		return 1
	}
	if w < h {
		return CountLatticePaths(h, w)
	}

	pathCount := CountLatticePaths(w-1, h) + CountLatticePaths(w, h-1)
	memoizeCache[Vertex{w, h}] = pathCount
	return pathCount
}

func Run() int {
	return CountLatticePaths(20, 20)
}
