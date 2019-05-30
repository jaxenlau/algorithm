package _1_max_area

// 给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
// 说明：你不能倾斜容器，且 n 的值至少为 2。
//
// 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
//
// 示例:
//
// 输入: [1,8,6,2,5,4,8,3,7]
// 输出: 49

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	area := 0
	for left < right {
		x := height[left]
		y := height[right]
		area = max(area, min(x, y)*(right-left))
		if x < y {
			left++
		} else {
			right--
		}
	}
	return area
}

func maxArea2(height []int) []int {
	left := 0
	right := len(height) - 1
	area := 0
	tempArea := 0
	post := [2]int{left, right}

	for left < right {
		x := height[left]
		y := height[right]

		if tempArea = min(x, y) * (right - left); tempArea > area {
			area = tempArea
			post[0] = left
			post[1] = right
		}
		if x < y {
			left++
		} else {
			right--
		}
	}
	return post[:]
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
