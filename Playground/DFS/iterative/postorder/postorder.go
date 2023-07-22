package postorder

type TreeNode[T any] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

// // binaryTree := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G'}
// func postOrder[T any](root *TreeNode[T]) (traversed []T) {

// }

// left -> right -> root
// expectedTraversal := []rune{'D', 'E', 'B', 'F', 'G', 'C', 'A'}
