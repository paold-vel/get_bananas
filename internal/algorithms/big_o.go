package algorithms

func SearchIndex(n []int, x int) []int { // O(n)
	indexes := []int{}
	for i := 0; i < len(n); i++ {
		if n[i] == x {
			indexes = append(indexes, i)
		}
	}
	return indexes
}
