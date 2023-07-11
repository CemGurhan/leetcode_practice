package stack

type Stack[T any] []T

func (s *Stack[T]) Push(item T) {
	*s = append(*s, item)
}

func (s *Stack[T]) Pop() T {
	poppedItem := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return poppedItem
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}
