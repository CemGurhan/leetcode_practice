package bfs

type TreeNode[T any] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

type Queue[T any] []*TreeNode[T]

func bfs[T any](root *TreeNode[T]) (levels [][]T) {
	nodesToProcess := Queue[T]{}
	if root != nil {
		nodesToProcess.Enqueue(root)
		levels = append(levels, []T{root.Val})
	} else {
		return levels
	}

	for !nodesToProcess.IsEmpty() {
		currentLevelNodes := []T{}
		currentLevelSize := len(nodesToProcess)
		for i := 0; i < currentLevelSize; i++ {
			currentItem := nodesToProcess.Dequeue()
			leftNode, rightNode := currentItem.Left, currentItem.Right
			if leftNode != nil {
				currentLevelNodes = append(currentLevelNodes, leftNode.Val)
				nodesToProcess.Enqueue(leftNode)
			}
			if rightNode != nil {
				currentLevelNodes = append(currentLevelNodes, rightNode.Val)
				nodesToProcess.Enqueue(rightNode)
			}
		}
		if len(currentLevelNodes) != 0 {
			levels = append(levels, currentLevelNodes)
		}
	}

	return levels
}

func (q *Queue[T]) Enqueue(val *TreeNode[T]) {
	*q = append(*q, val)
}

func (q *Queue[T]) Dequeue() *TreeNode[T] {
	if q.IsEmpty() {
		return nil
	}
	dequeuedNode := (*q)[0]
	*q = (*q)[1:]
	return dequeuedNode
}

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}
