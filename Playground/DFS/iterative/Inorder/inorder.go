package inorder

import utils "github.com/cemgurhan/leetcodepractice/Playground/DFS/iterative/utils"

type TreeNode[T any] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func dfs[T any](root *TreeNode[T]) (traversed []T) {
	stack := make(utils.Stack[*TreeNode[T]], 0)
	curr := root

	for curr != nil || !stack.IsEmpty() {
		for curr != nil {
			stack.Push(curr)
			curr = curr.Left
		}
		curr = stack.Pop()
		traversed = append(traversed, curr.Val)
		curr = curr.Right
	}
	return
}

// binaryTree := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G'}
// expectedTraversal := []rune{'D', 'B', 'E', 'A', 'F', 'C', 'G'}

// [A, B, D]
