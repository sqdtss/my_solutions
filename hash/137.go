package solution

// 137. 只出现一次的数字 II
// 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。

// 示例 1：
// 输入：nums = [2,2,3,2]
// 输出：3

// 示例 2：
// 输入：nums = [0,1,0,1,0,1,99]
// 输出：99

// 提示：
// 1 <= nums.length <= 3 * 104
// -231 <= nums[i] <= 231 - 1
// nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次

// 方法1: 位运算
// golang中int和机器有关，64位机器中int占8个字节（和int64一样）
func singleNumber(nums []int) int {
	var count [32]int32
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			count[i] += int32(num & 1)
			num >>= 1
		}
	}
	ans := int32(0)
	for i := 0; i < 32; i++ {
		ans <<= 1
		ans |= count[31-i] % 3
	}
	return int(ans)
}

// 方法2: hash

// func singleNumber(nums []int) int {
//     freq := map[int]int{}
//     for _, v := range nums {
//         freq[v]++
//     }
//     for num, occ := range freq {
//         if occ == 1 {
//             return num
//         }
//     }
//     return 0 // 不会发生，数据保证有一个元素仅出现一次
// }
