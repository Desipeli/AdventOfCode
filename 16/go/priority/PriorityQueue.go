package priority

import (
	"fmt"
)

type PriorityQueueNode struct {
	Key   int
	Value interface{}
}

type PrioirityQueue struct {
	tree []PriorityQueueNode
}

func NewPriorityQueue() *PrioirityQueue {
	return &PrioirityQueue{tree: []PriorityQueueNode{}}
}

func (pq *PrioirityQueue) Insert(node PriorityQueueNode) {
	pq.tree = append(pq.tree, node)
	pq.pushElementUp(len(pq.tree) - 1)
}

func (pq *PrioirityQueue) Pop() (PriorityQueueNode, bool) {
	if len(pq.tree) == 0 {
		return PriorityQueueNode{}, false
	}
	smallest := pq.tree[0]
	pq.tree[0] = pq.tree[len(pq.tree)-1]
	pq.tree = pq.tree[:len(pq.tree)-1]
	pq.pushElementDown(0)
	return smallest, true
}

func (pq *PrioirityQueue) pushElementUp(index int) {
	if index == 0 {
		return
	}
	parentIndex := (index - 1) / 2
	parent := pq.tree[parentIndex]
	current := pq.tree[index]
	if parent.Key > current.Key {
		pq.tree[parentIndex] = current
		pq.tree[index] = parent
		pq.pushElementUp(parentIndex)
	}
}

func (pq *PrioirityQueue) pushElementDown(index int) {
	size := len(pq.tree)
	smallest := index

	leftChild := 2*index + 1
	rightChild := 2*index + 2

	if leftChild < size && pq.tree[leftChild].Key < pq.tree[smallest].Key {
		smallest = leftChild
	}
	if rightChild < size && pq.tree[rightChild].Key < pq.tree[smallest].Key {
		smallest = rightChild
	}

	if smallest != index {
		pq.tree[index], pq.tree[smallest] = pq.tree[smallest], pq.tree[index]
		pq.pushElementDown(smallest)
	}
}

func (pq *PrioirityQueue) Display() {
	fmt.Println(pq.tree)
}
