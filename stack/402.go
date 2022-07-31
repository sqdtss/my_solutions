package solution

import "strings"

/* my solution one */
// func removeKdigits(num string, k int) string {
// 	if len(num) <= k {
// 		return "0"
// 	}
// 	ans := ""
// 	for i := range num {
// 		for k > 0 && len(ans) > 0 && ans[len(ans)-1] > num[i] {
// 			ans = ans[:len(ans)-1]
// 			k--
// 		}
// 		ans = ans + string(num[i])
// 	}
// 	if k > 0 {
// 		ans = ans[:len(ans)-k]
// 	}

// 	// 去除前导0，注意ans=“0”的情况
// 	for {
// 		if len(ans) > 1 && ans[0] == '0' {
// 			ans = ans[1:]
// 		} else {
// 			break
// 		}
// 	}
// 	return ans
// }

func removeKdigits(num string, k int) string {
	stack := []byte{}
	for i := range num {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	if k > 0 {
		stack = stack[:len(stack)-k]
	}

	// 去除前导0，注意ans=“0”的情况
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}
