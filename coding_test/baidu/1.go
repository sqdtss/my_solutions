package main

import "fmt"

func main() {
	var t int
	for fmt.Scan(&t); t > 0; t-- {
		var in string
		fmt.Scan(&in)
		flag := false
		if len(in) == len("Baidu") {
			var counts [5]int
			for _, c := range in {
				if c == 'B' {
					counts[0]++
				} else if c == 'a' {
					counts[1]++
				} else if c == 'i' {
					counts[2]++
				} else if c == 'd' {
					counts[3]++
				} else if c == 'u' {
					counts[4]++
				}
			}
			flag = true
			for i := 0; i < 5; i++ {
				if counts[i] != 1 {
					flag = false
				}
			}
		}
		if flag {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
