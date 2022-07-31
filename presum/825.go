package presum

import "sort"

//825. 适龄的朋友
//在社交媒体网站上有 n 个用户。给你一个整数数组 ages ，其中 ages[i] 是第 i 个用户的年龄。
//如果下述任意一个条件为真，那么用户 x 将不会向用户 y（x != y）发送好友请求：
//ages[y] <= 0.5 * ages[x] + 7
//ages[y] > ages[x]
//ages[y] > 100 && ages[x] < 100
//否则，x 将会向 y 发送一条好友请求。
//注意，如果 x 向 y 发送一条好友请求，y 不必也向 x 发送一条好友请求。另外，用户不会向自己发送好友请求。
//返回在该社交媒体网站上产生的好友请求总数。
//
//示例 1：
//输入：ages = [16,16]
//输出：2
//解释：2 人互发好友请求。
//
//示例 2：
//输入：ages = [16,17,18]
//输出：2
//解释：产生的好友请求为 17 -> 16 ，18 -> 17 。
//
//示例 3：
//输入：ages = [20,30,100,110,120]
//输出：3
//解释：产生的好友请求为 110 -> 100 ，120 -> 110 ，120 -> 100 。
//
//提示：
//n == ages.length
//1 <= n <= 2 * 104
//1 <= ages[i] <= 120

func numFriendRequests(ages []int) (ans int) {
	sort.Ints(ages)
	const maxAge = 121
	var cnt, pre [maxAge]int
	for _, age := range ages {
		cnt[age]++
	}
	for i := 1; i < maxAge; i++ {
		pre[i] = pre[i-1] + cnt[i]
	}
	for i := 15; i < maxAge; i++ {
		if cnt[i] > 0 {
			bound := i/2 + 7
			ans += cnt[i] * (pre[i] - pre[bound] - 1) // 这里每个age为i的除了自己其他的均可交朋友，及-1
		}
	}
	return
}
