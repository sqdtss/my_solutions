package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n, &q)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < q; i++ {
		var op, i, x int
		fmt.Scan(&op, &i, &x)
		if op == 1 {
			a[i-1] = x
		} else {
			ans := 0
			for idx := 0; idx < i; idx++ {
				if a[idx] == x {
					ans++
				}
			}
			fmt.Println(ans)
		}
	}
}
