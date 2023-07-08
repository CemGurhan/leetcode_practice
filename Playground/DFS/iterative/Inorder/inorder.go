package inorder

type Stack[T any] []T

type TreeNode[T any] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func dfs[T any](root *TreeNode[T]) (traversed []T) {
	stack := make(Stack[*TreeNode[T]], 0)
	curr := root

	for curr != nil || !stack.isEmpty() {
		for curr != nil {
			stack.push(curr)
			curr = curr.Left
		}
		curr = stack.pop()
		traversed = append(traversed, curr.Val)
		curr = curr.Right
	}
	return
}

// stack := [A]
// tree := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G'}
// expectedTraversal := []rune{'D', 'B', 'E', 'A', 'F', 'C', 'G'}

func (s *Stack[T]) push(item T) {
	*s = append(*s, item)
}

func (s *Stack[T]) pop() T {
	poppedItem := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return poppedItem
}

func (s *Stack[T]) isEmpty() bool {
	return len(*s) == 0
}

// [A,B,C]
