package threeSumClosest

import (
	"sort"
)

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和
。假定每组输入只存在唯一答案。


示例：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。


提示：

3 <= nums.length <= 10^3
-10^3 <= nums[i] <= 10^3
-10^4 <= target <= 10^4

Related Topics 数组 双指针 排序
👍 821 👎 0
*/

// leetcode submit region begin(Prohibit modification and deletion)
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var (
		ret = 10000
		n   = len(nums)
	)

	for a := 0; a < n; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}

		b := a + 1
		c := n - 1

		for b < c {
			temp := nums[a] + nums[b] + nums[c]
			if temp == target {
				return temp
			}

			if distance(ret, target) > distance(temp, target) {
				ret = temp
			}

			if temp < target {
				b++
				for b < c && nums[b] == nums[b-1] {
					b++
				}
			} else {
				c--
				for c > b && nums[c] == nums[c+1] {
					c--
				}
			}
		}
	}

	return ret
}

func distance(x, y int) int {
	return abs(x - y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// leetcode submit region end(Prohibit modification and deletion)
