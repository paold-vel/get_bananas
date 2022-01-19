package main

import (
	"fmt"
	"get_bananas/internal/algorithms"
)

func main() {
	n := []int{7, 2, 7, 4, 7, 6, 7}
	indexes := algorithms.SearchIndex(n, 7)
	for i := 0; i < len(indexes); i++ {
		fmt.Println(indexes[i])
	}
}
