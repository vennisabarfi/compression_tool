package main

// Node struct
type Node struct {
	char  rune
	count int
	left  *Node
	right *Node
}

// https://reintech.io/blog/comprehensive-guide-huffman-coding-algorithm-go

// implementing algo by creating a priority queue. Ordering queue according to the frequency of the characters
type PriorityQueue []*Node

// length of queue
func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].count < pq[j].count
}

// swap
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// push on to heap
func (pq *PriorityQueue) Push(x interface{}) {

	*pq = append(*pq, x.(*Node))
}

// pop from heap
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
