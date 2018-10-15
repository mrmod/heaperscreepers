package minheap

import "fmt"

// HeapObject Stores a something
type HeapObject struct {
	Index, Score int
	Value        interface{}
}

// Heap A collection of HeapObject
type Heap []HeapObject

// ROOT is the root node
const ROOT = 1

// LessThanParent True if ho.Index/2 Score less than our score
func (heap Heap) LessThanParent(ho HeapObject) bool {
	if ho.Index == 1 {
		return false
	}

	if parent := heap[ho.Index/2]; ho.Score < parent.Score {
		return true
	}
	return false
}

// Swap Swap a into b's position and b to a's
func (heap Heap) Swap(child, parent HeapObject) Heap {
	fmt.Printf("Swapping %v with %v\n", child, parent)
	t := heap[parent.Index]
	t.Index = child.Index

	heap[child.Index] = t
	heap[parent.Index] = child
	child.Index = parent.Index
	return heap
}

// Insert Add a new member
func (heap Heap) Insert(ho HeapObject) Heap {
	ho.Index = len(heap)
	newHeap := append(heap, ho)
	if ho.Index == 1 {
		return newHeap
	}

	for newHeap.LessThanParent(ho) {
		newHeap = newHeap.Swap(ho, newHeap[ho.Index/2])
		ho = newHeap[ho.Index/2]
	}
	return newHeap
}

// Pop Return the Root
func (heap Heap) Pop() (Heap, HeapObject) {
	root := heap[1]
	newHeap := NewMinHeap()

	newRoot := heap[len(heap)-1]
	newRoot.Index = 1
	// Root the tree at the newRoot
	newHeap = append(newHeap, newRoot)
	newHeap = append(newHeap, heap[2:]...)
	// While the moved element Score is > than at least one child
	for !newHeap.LessThanIndex(newRoot.Index*2, (newRoot.Index*2)+1) {
		// Swap newRoot with the lesser child
		// TODO: :Complete
	}

	}
}

// Peek Show the root
func (heap Heap) Peek() HeapObject {
	return heap[ROOT]
}

// NewMinHeap Min is always at the head
func NewMinHeap() Heap {
	return Heap{HeapObject{}}
}
