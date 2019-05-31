package _42_trapping_tain_water

import (
	"github.com/owenliu1122/algorithm/algorithm/goalng/ds/stack"
)

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
// 上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。
//
// 示例:
//
// 输入: [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出: 6

// 解法 1： 暴力， 遍历每一个元素，并且向左和向右找最高的
// 时间复杂度 O(n^2)
// 空间复杂度 O(1)
func trapViolent(height []int) int {
	count := len(height)
	rainfall := 0
	for i := 1; i < count-1; i++ {
		leftMax := height[i]
		rightMax := height[i]

		// 左侧最大
		for l := i - 1; l >= 0; l-- {
			if leftMax < height[l] {
				leftMax = height[l]
			}
		}
		// 右侧最大
		for r := i + 1; r < count; r++ {
			if rightMax < height[r] {
				rightMax = height[r]
			}
		}
		rainfall += min(leftMax, rightMax) - height[i]
	}
	return rainfall
}

// 解法 2： 动态编程 Dynamic Programming，先记录每个点的左右最高点
// 然后计算每个点的接收的雨量，需借助额外空间，
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func trapDynamic(height []int) int {
	count := len(height)
	if count == 0 {
		return 0
	}

	// 左侧最大集合
	leftMax := make([]int, count)
	leftMax[0] = height[0]
	for l := 1; l < count; l++ {
		leftMax[l] = max(leftMax[l-1], height[l])
	}

	// 右侧最大集合
	rightMax := make([]int, count)
	rightMax[count-1] = height[count-1]
	for r := count - 2; r >= 0; r-- {
		rightMax[r] = max(rightMax[r+1], height[r])
	}

	// 计算结果
	rainfall := 0
	for i := 1; i < count-1; i++ {
		rainfall += min(leftMax[i], rightMax[i]) - height[i]
	}
	return rainfall
}

// 解法 3： 栈, 分层思维
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func trapStack(height []int) int {
	count := len(height)
	// 计算结果
	rainfall := 0
	s := stack.New(count)

	for i := 0; i < count; i++ {
		for !s.IsEmpty() {
			top := (*s.Top()).(int)
			if height[i] <= height[top] {
				break
			}
			s.Pop()
			if s.IsEmpty() {
				break
			}

			nexTop := (*s.Top()).(int)

			// 存水两端柱子之间的距离
			dis := i - nexTop - 1

			// 当前柱子和栈顶索引的柱子最小值到上一个比当前栈顶低的柱子之间的差值，也就是能存水的层高
			h := min(height[i], height[nexTop]) - height[top]

			// 累加一层
			rainfall += dis * h
		}
		s.Push(i)
	}
	return rainfall
}

// 解法 4： 双指针
// 时间复杂度 O(n)
// 空间复杂度 O(1)
func trapTwoPointers(height []int) int {
	var (
		rainfall = 0
		left     = 0
		right    = len(height) - 1
		leftMax  = 0
		rightMax = 0
	)

	for left < right {
		if height[left] < height[right] {
			if leftMax < height[left] {
				leftMax = height[left]
			} else {
				rainfall += leftMax - height[left]
			}
			left++
		} else {
			if rightMax < height[right] {
				rightMax = height[right]
			} else {
				rainfall += rightMax - height[right]
			}
			right--
		}
	}
	return rainfall
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
