package recursive

type TreeNode[T any] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func inOrder[T any](root *TreeNode[T], traversed *[]T) {
	if root == nil {
		return
	}

	inOrder(root.Left, traversed)
	*traversed = append(*traversed, root.Val)
	inOrder(root.Right, traversed)

	return
}
