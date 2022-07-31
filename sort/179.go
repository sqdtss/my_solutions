package sort

import (
	"sort"
	"strconv"
	"strings"
)

//179. 最大数
//给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
//注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
//
//示例 1：
//输入：nums = [10,2]
//输出："210"
//
//示例 2：
//输入：nums = [3,30,34,5,9]
//输出："9534330"
//
//提示：
//1 <= nums.length <= 100
//0 <= nums[i] <= 109

func largestNumber(nums []int) string {
	var numsStringSz []string
	for _, num := range nums {
		numsStringSz = append(numsStringSz, strconv.Itoa(num))
	}
	sort.Slice(numsStringSz, func(i, j int) bool {
		tmp1, tmp2 := numsStringSz[i]+numsStringSz[j], numsStringSz[j]+numsStringSz[i]
		return tmp1 > tmp2
	})
	resSB := strings.Builder{}
	for _, str := range numsStringSz {
		resSB.WriteString(str)
	}
	res := resSB.String()
	for 0 < len(res)-1 && res[0] == '0' {
		res = res[1:]
	}
	return res
}
