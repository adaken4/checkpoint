package checkpoint

import "fmt"

func Chunk(slice []int, size int) {
	var Slice [][]int
	var chunk []int
	if size == 0 {
		fmt.Println()
		return
	}
	for _, v := range slice {
		chunk = append(chunk, v)
		if size == len(chunk) {
			Slice = append(Slice, chunk)
			chunk = []int{}
		}
	}
	if len(chunk) > 0 {
		Slice = append(Slice, chunk)
	}
	fmt.Println(Slice)
}
