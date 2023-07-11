package postorder

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// binaryTree := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G'}
func createBinaryTree[T rune]() TreeNode[T] {
	return TreeNode[T]{
		Val: 'A',
		Left: &TreeNode[T]{
			Val: 'B',
			Left: &TreeNode[T]{
				Val:   'D',
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode[T]{
				Val:   'E',
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode[T]{
			Val: 'C',
			Left: &TreeNode[T]{
				Val:   'F',
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode[T]{
				Val:   'G',
				Left:  nil,
				Right: nil,
			},
		},
	}
}

func TestDFS_WithBinaryTree_ReturnsValidPostOrderTraversal(t *testing.T) {
	binaryTree := createBinaryTree[rune]()
	actualTraversal := postOrder[rune](&binaryTree)
	expectedTraversal := []rune{'D', 'E', 'B', 'F', 'G', 'C', 'A'}

	require.Equal(t, expectedTraversal, actualTraversal)
}
