package maxdepthofbinarytree

import m "github.com/cemgurhan/leetcodepractice/MaxDepthOfBinaryTree/models"

type Queue []*m.TreeNode

func maxDepth(root *m.TreeNode) int {
	queue := make(Queue, 0)
	queue.Enqueue(root)
	maxDepth := 0

	for !queue.IsEmpty() {
		maxDepth++
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			dequeuedNode := queue.Dequeue()
			queue.Enqueue(dequeuedNode.Left)
			queue.Enqueue(dequeuedNode.Right)
		}
	}

	return maxDepth
}

func (q *Queue) Enqueue(node *m.TreeNode) {
	if node == nil {
		return
	}
	*q = append(*q, node)
}

func (q *Queue) Dequeue() *m.TreeNode {
	dequeuedNode := (*q)[0]
	*q = (*q)[1:]
	return dequeuedNode
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
