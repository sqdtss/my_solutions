package solution

// 633. 平方数之和
// 给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c 。

// 示例 1：
// 输入：c = 5
// 输出：true
// 解释：1 * 1 + 2 * 2 = 5

// 示例 2：
// 输入：c = 3
// 输出：false

// 示例 3：
// 输入：c = 4
// 输出：true

// 示例 4：
// 输入：c = 2
// 输出：true

// 示例 5：
// 输入：c = 1
// 输出：true

// 提示：
// 0 <= c <= 231 - 1

//方法1: hash表做法
// func judgeSquareSum(c int) bool {
// 	mp := make(map[int]bool)
// 	for i := 0; i*i <= c; i++ {
// 		mp[i*i] = true
// 	}
// 	for i := 0; i*i <= c; i++ {
// 		if mp[c-i*i] {
// 			return true
// 		}
// 	}
// 	return false
// }

//方法2：双指针
// func judgeSquareSum(c int) bool {
//     left, right := 0, int(math.Sqrt(float64(c)))
//     for left <= right {
//         sum := left*left + right*right
//         if sum == c {
//             return true
//         } else if sum > c {
//             right--
//         } else {
//             left++
//         }
//     }
//     return false
// }

//方法3：费马平方和定理
// 一个非负整数 c 如果能够表示为两个整数的平方和，当且仅当 c 的所有形如 4k + 3 的质因子的幂均为偶数。
func judgeSquareSum(c int) bool {
	for base := 2; base*base <= c; base++ {
		// 如果不是因子，枚举下一个
		if c%base > 0 {
			continue
		}

		// 计算 base 的幂
		exp := 0
		for ; c%base == 0; c /= base {
			exp++
		}

		// 根据 Sum of two squares theorem 验证
		if base%4 == 3 && exp%2 != 0 {
			return false
		}
	}

	// 例如 11 这样的用例，由于上面的 for 循环里 base * base <= c ，base == 11 的时候不会进入循环体
	// 因此在退出循环以后需要再做一次判断
	return c%4 != 3
}
