package btreelevelorder

// [102] 给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
//
// 例如:
// 给定二叉树: [3,9,20,null,null,15,7],
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回其层次遍历结果：
//
// [
//  [3],
//  [9,20],
//  [15,7]
// ]

type Stack struct {
	Size int
	Buff []*TreeNode
}

func (s *Stack) IsEmpty() bool {
	return len(s.Buff) == 0
}

func (s *Stack) IsFull() bool {
	return len(s.Buff) == s.Size
}

func (s *Stack) Push(node *TreeNode) {
	s.Buff = append(s.Buff, node)
}

func (s *Stack) Pop() *TreeNode {
	topIdx := len(s.Buff) - 1
	ret := s.Buff[topIdx]
	s.Buff = s.Buff[:topIdx]
	return ret
}

type Quene struct {
	Size int
	Buff []*TreeNode
}

func (q *Quene) IsEmpty() bool {
	return len(q.Buff) == 0
}

func (q *Quene) IsFull() bool {
	return len(q.Buff) == q.Size
}

func (q *Quene) Push(node *TreeNode) {
	q.Buff = append(q.Buff, node)
}

func (q *Quene) Pop() *TreeNode {
	ret := q.Buff[0]
	q.Buff = q.Buff[1:]
	return ret
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	subNodes := []int{}
	var curLvCnt int = 1
	var nextLvCnt int = 0

	q := Quene{50, []*TreeNode{}}

	if nil != root {
		q.Push(root)
	}

	for !q.IsEmpty() {

		p := q.Pop()

		curLvCnt--

		subNodes = append(subNodes, p.Val)

		if nil != p.Left {
			q.Push(p.Left)
			nextLvCnt++
		}

		if nil != p.Right {
			q.Push(p.Right)
			nextLvCnt++
		}

		if 0 == curLvCnt {
			curLvCnt = nextLvCnt
			nextLvCnt = 0
			result = append(result, subNodes)
			subNodes = append([]int{})
		}

	}

	return result
}
