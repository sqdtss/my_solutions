package solution

//739. 每日温度
//请根据每日 气温 列表 temperatures ，请计算在每一天需要等几天才会有更高的温度。
//如果气温在这之后都不会升高，请在该位置用 0 来代替。
//
//示例 1:
//输入: temperatures = [73,74,75,71,69,72,76,73]
//输出: [1,1,4,2,1,1,0,0]
//
//示例 2:
//输入: temperatures = [30,40,50,60]
//输出: [1,1,1,0]
//
//示例 3:
//输入: temperatures = [30,60,90]
//输出: [1,1,0]
//
//提示：
//1 <= temperatures.length <= 105
//30 <= temperatures[i] <= 100

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures)) // ans初始化就是0，没变的就保持0了！
	stack := []int{}
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return ans
}
