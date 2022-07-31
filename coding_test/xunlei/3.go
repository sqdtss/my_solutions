package main

import "fmt"

func main() {
	var nums []int
	var num int
	sum := 0
	for _, err := fmt.Scan(&num); err == nil; _, err = fmt.Scan(&num) {
		nums = append(nums, num)
		sum += num
	}
	ans := 0
	avg := sum / len(nums)
	for i, num := range nums {
		if num < avg {
			nums[i+1] -= avg - num
			ans++
		} else if num > avg {
			nums[i+1] += num - avg
			ans++
		}
	}
	fmt.Println(ans)
}
