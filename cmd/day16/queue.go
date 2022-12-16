package main

import "container/heap"

// Queue is a priority queue of Valves.  After items are pushed on to
// the queue, they will be popped off with largest priority first.
type Queue []*Node

// Node combines a ValveID with its priority.
type Node struct {
	Value    ValveID
	Priority int
	index    int
}

// compile-time interface check.
// The priority queue is implemented as a heap.
var _ heap.Interface = new(Queue)

// Len implements heap.Interface.
func (pq Queue) Len() int { return len(pq) }

// Less implements heap.Interface using the (negative) priority of the item.
func (pq Queue) Less(i, j int) bool {
	// we want Pop to give the highest priority, not lowest, so we use greater
	return pq[i].Priority > pq[j].Priority
}

// Swap implements heap.Interface.
func (pq Queue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push implements heap.Interface.
// Do not use this method to push an item on to the queue. Instead, use heap.Push().
func (pq *Queue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}

// Pop implements heap.Interface.
// Do not use this method to pop an item off the queue. Instead, use heap.Pop().
func (pq *Queue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[:n-1]
	return item
}

// Update modifies the priority and value of an item in the queue.
func (pq *Queue) Update(item *Node, v ValveID, priority int) {
	item.Value = v
	item.Priority = priority
	heap.Fix(pq, item.index)
}
