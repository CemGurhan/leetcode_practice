package maxdepthofbinarytree

import m "github.com/cemgurhan/leetcodepractice/MaxDepthOfBinaryTree/models"

func maxDepth(root *m.TreeNode) int {
	if root == nil {
		return 0
	}

	leftCount := maxDepth(root.Left)
	rightCount := maxDepth(root.Right)

	return 1 + maxInt(leftCount, rightCount)
}

func maxInt(left, right int) int {
	if left > right {
		return left
	}

	return right
}
