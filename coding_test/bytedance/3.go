package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		minute := int(math.Abs(float64(x - y)))
		ans := 0
		nowSum := 0
		step := 1
		for nowSum < minute {
			nowSum += step
			ans++
			if nowSum < minute {
				nowSum += step
			} else {
				break
			}
			ans++
			step++
		}
		fmt.Println(ans)
	}
}
