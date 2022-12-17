package prioqueue

import "container/heap"

// Queue is a generic priority queue.  After items are pushed on to
// the queue, they will be popped off with largest priority first.
//
// The zero value is ready to use:
//    q := new(Queue)
//
// An existing slice can be converted to a priority Queue:
//    nodes := []*Node[int] {
//        {Value: 1, Priority: 1},
//        {Value: 2, Priority: 2},
//        {Value: 3, Priority: 3}}
// 	  q := Queue[int](nodes)
//    heap.Init(&q)
//
// Use heap.Push() and heap.Pop() to push and pop items on to the queue:
//    heap.Push(q, &Node{ Value: twod.Point{X: 1, Y: 2}, Priority: 3})
//    node, ok := heap.Pop(q).(*Node)
//
// Keep a pointer to a node in the queue, and then use Update()
// to update the value and/or priority of the node:
//    heap.Push(q, node)
//    q.Update(node, node.Value, 4) // keeps the existing value
//
// adapted from the example at: https://pkg.go.dev/container/heap
type Queue[T any] []*Node[T]

// Node combines a value with its priority.
type Node[T any] struct {
	Value    T
	Priority int
	index    int
}

// Len implements heap.Interface.
func (pq Queue[T]) Len() int { return len(pq) }

// Less implements heap.Interface using the (negative) priority of the item.
func (pq Queue[T]) Less(i, j int) bool {
	// we want Pop to give the highest priority, not lowest, so we use greater
	return pq[i].Priority > pq[j].Priority
}

// Swap implements heap.Interface.
func (pq Queue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push implements heap.Interface.
// Do not use this method to push an item on to the queue. Instead, use heap.Push().
func (pq *Queue[T]) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node[T])
	item.index = n
	*pq = append(*pq, item)
}

// Pop implements heap.Interface.
// Do not use this method to pop an item off the queue. Instead, use heap.Pop().
func (pq *Queue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[:n-1]
	return item
}

// Update modifies the priority and value of an item in the queue.
func (pq *Queue[T]) Update(item *Node[T], v T, priority int) {
	item.Value = v
	item.Priority = priority
	heap.Fix(pq, item.index)
}
