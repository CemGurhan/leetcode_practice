package bfs

import (
	"reflect"
	"testing"
)

// binaryTree := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G'}
func createBinaryTree() TreeNode[rune] {
	return TreeNode[rune]{
		Val: 'A',
		Left: &TreeNode[rune]{
			Val: 'B',
			Left: &TreeNode[rune]{
				Val:   'D',
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode[rune]{
				Val:   'E',
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode[rune]{
			Val: 'C',
			Left: &TreeNode[rune]{
				Val:   'F',
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode[rune]{
				Val:   'G',
				Left:  nil,
				Right: nil,
			},
		},
	}
}

func TestBFS(t *testing.T) {
	binaryTree := createBinaryTree()
	expectedLevels := [][]rune{{'A'}, {'B', 'C'}, {'D', 'E', 'F', 'G'}}

	actualLevels := bfs(&binaryTree)

	if ok := reflect.DeepEqual(expectedLevels, actualLevels); !ok {
		t.Fatalf("want %c, got %c", expectedLevels, actualLevels)
	}
}
