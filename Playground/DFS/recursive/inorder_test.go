package recursive

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

func TestDFS_WithValidBinaryTree_ReturnsCorrectInOrderTraversal(t *testing.T) {
	binaryTree := createBinaryTree[rune]()
	traversed := make([]rune, 0)
	inOrder[rune](&binaryTree, &traversed)
	expectedTraversal := []rune{'D', 'B', 'E', 'A', 'F', 'C', 'G'}

	require.Equal(t, expectedTraversal, traversed)
}
