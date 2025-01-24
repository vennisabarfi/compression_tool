package main

// huffmantree struct
type huffmanTree struct {
	nodes    []huffmanNode
	nextNode int
}

// struct to hold a node in the tree
type huffmanNode struct {
	left, right           uint16
	leftValue, rightValue uint16
}
