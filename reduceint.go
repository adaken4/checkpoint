package checkpoint

import "fmt"

func ReduceInt(a []int, f func(int, int) int) {
	rslt := a[0]
	for i := 1; i < len(a); i++ {
		rslt = f(rslt, a[i])
	}
	fmt.Println(rslt)
}
