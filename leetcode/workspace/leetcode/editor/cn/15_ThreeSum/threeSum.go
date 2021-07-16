package threeSum

import (
	"sort"
)

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
复的三元组。

注意：答案中不可以包含重复的三元组。


示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]


示例 2：

输入：nums = []
输出：[]


示例 3：

输入：nums = [0]
输出：[]


提示：

0 <= nums.length <= 3000
-105 <= nums[i] <= 105

Related Topics 数组 双指针 排序
👍 3498 👎 0*/

// leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	ret := make([][]int, 0, 10)

	for a := 0; a < n; a++ {

		if a > 0 && nums[a] == nums[a-1] { // a 要判断大于零，不然会丢失答案
			continue
		}

		c := n - 1
		target := -(nums[a])

		for b := a + 1; b < n; b++ {

			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}

			for b < c && nums[b]+nums[c] > target {
				c--
			}

			if b == c {
				break
			}

			if nums[b]+nums[c] == target {
				ret = append(ret, []int{nums[a], nums[b], nums[c]})
			}
		}
	}

	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
