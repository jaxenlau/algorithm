package validprentheses

// [20] 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
//
// 注意空字符串可被认为是有效字符串。
//
// 示例 1:
//
// 输入: "()"
// 输出: true
//
//
// 示例 2:
//
// 输入: "()[]{}"
// 输出: true
//
//
// 示例 3:
//
// 输入: "(]"
// 输出: false
//
//
// 示例 4:
//
// 输入: "([)]"
// 输出: false
//
//
// 示例 5:
//
// 输入: "{[]}"
// 输出: true
//

type Stack struct {
	Size int
	Buff []rune
}

func (s *Stack) IsEmpty() bool {
	return len(s.Buff) == 0
}

func (s *Stack) IsFull() bool {
	return len(s.Buff) == s.Size
}

func (s *Stack) Push(node rune) {
	s.Buff = append(s.Buff, node)
}

func (s *Stack) Pop() rune {
	topIdx := len(s.Buff) - 1
	ret := s.Buff[topIdx]
	s.Buff = s.Buff[:topIdx]
	return ret
}

func isValid(s string) bool {
	cmap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := &Stack{
		len(s),
		make([]rune, 0, len(s)),
	}

	for _, c := range s {
		l, ok := cmap[c]
		if !ok {
			stack.Push(c)

		} else if stack.IsEmpty() || l != stack.Pop() {
			return false
		}
	}
	return stack.IsEmpty()
}
