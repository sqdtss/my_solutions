package solution

import "sort"

// hash

func groupAnagrams(strs []string) (ans [][]string) {
	mp := make(map[string][]string)
	for _, str := range strs {
		tmp := []byte(str)
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		sortedStr := string(tmp)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
