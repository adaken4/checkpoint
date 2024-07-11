package checkpoint

import "os"

func FoldInt(f func(int, int) int, a []int, n int) {
	result := n
	for i := 0; i < len(a); i++ {
		result = f(result, a[i])
	}
	os.Stdout.WriteString(Itoa(result) + "\n")
}
