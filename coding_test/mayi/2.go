package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	var k int
	fmt.Scan(&str)
	fmt.Scan(&k)
	ans := 0
	for i := 0; i < len(str); i++ {
		for j := i + 1; j <= len(str); j++ {
			num, _ := strconv.Atoi(str[i:j])
			if num < k {
				ans++
			} else {
				break
			}
		}
	}
	fmt.Println(ans)
}
