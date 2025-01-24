package main

import (
	"container/heap"
	"fmt"
	"log"
	"os"
)

func frequencyCounter(filename string) (frequency map[rune]int, pq *PriorityQueue) {
	buffer, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	// convert buffer into string for counter
	s := string(buffer)
	// frequency counter that creates a map of runes and integers
	frequency = make(map[rune]int)
	// iterate over string and create count
	for _, char := range s {
		frequency[char] = frequency[char] + 1
	}

	// build tree here
	// initialize a heap
	pq = &PriorityQueue{}
	heap.Init(pq)

	// push char and count onto heap
	for char, count := range frequency {
		heap.Push(pq, &Node{char: char, count: count})
	}

	fmt.Println("Here's the priority queue", pq)

	// return sorted frequency? work on this.
	// return frequency map
	return frequency, pq

}

// build tree from priority queue from frequency counter
func BuildTree(pq *PriorityQueue) *Node {

	// while the priority queue exists build tree and push parent node/new tree onto heap
	for pq.Len() > 1 {
		first := heap.Pop(pq).(*Node)
		second := heap.Pop(pq).(*Node)

		parent := &Node{
			count: first.count + second.count,
			left:  first,
			right: second,
		}
		heap.Push(pq, parent)
	}
	return heap.Pop(pq).(*Node)
}

func FileCompressor(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	// check for error to be returned and exit if file not found
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	// extract frequency using frequency counter function
	frequency, _ := frequencyCounter(filename)
	for char, count := range frequency {
		fmt.Printf("%c: %d\n", char, count)
	}

}
