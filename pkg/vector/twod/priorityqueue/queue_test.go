package priorityqueue

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nealmcc/aoc2022/pkg/vector/twod"
)

func TestQueue_pushAndLen(t *testing.T) {
	q := make(Queue, 0)

	nodes := []*Node{
		{
			Value:    twod.Point{X: 4, Y: 2},
			Priority: 42,
		},
		{
			Value:    twod.Point{X: 6, Y: 7},
			Priority: 67,
		},
		{
			Value:    twod.Point{X: 0, Y: 0},
			Priority: 0,
		},
	}

	for _, n := range nodes {
		q.Push(n)
	}

	require.Equal(t, 3, q.Len())
}

func TestQueue_initAndPop(t *testing.T) {
	nodes := []*Node{
		{
			Value:    twod.Point{X: 4, Y: 2},
			Priority: 42,
		},
		{
			Value:    twod.Point{X: 6, Y: 7},
			Priority: 67,
		},
		{
			Value:    twod.Point{X: 0, Y: 0},
			Priority: 0,
		},
	}

	q := Queue(nodes)
	fmt.Println(q)

	heap.Init(&q)
	fmt.Println(q)

	node, ok := heap.Pop(&q).(*Node)
	require.True(t, ok)
	assert.Equal(t, twod.Point{X: 6, Y: 7}, node.Value)
	assert.Equal(t, 67, node.Priority)

	node, ok = heap.Pop(&q).(*Node)
	require.True(t, ok)
	assert.Equal(t, twod.Point{X: 4, Y: 2}, node.Value)
	assert.Equal(t, 42, node.Priority)
}

func ExampleQueue_Update() {
	items := map[twod.Point]int{
		{X: 4, Y: 2}: 42,
		{X: 6, Y: 7}: 67,
		{X: 0, Y: 0}: 0,
	}

	q := new(Queue)
	for val, prio := range items {
		heap.Push(q, &Node{Value: val, Priority: prio})
	}

	// push the item on to the queue (it will have priority 0)
	newItem := &Node{Value: twod.Point{X: 99, Y: 99}}
	heap.Push(q, newItem)

	// now update the item's priority:
	q.Update(newItem, newItem.Value, 9999)

	node, _ := heap.Pop(q).(*Node)
	fmt.Printf("priority: %2d value: %+v\n", node.Priority, node.Value)
	// Output: priority: 9999 value: {X:99 Y:99}
}
