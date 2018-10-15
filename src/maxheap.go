package heap

import "fmt"

// MaxHeap
type MaxHeap []HeapObject

// Peek Return the head
func (heap MaxHeap) Peek() HeapObject {
	return heap[1]
}

// Swap two elements
func (heap MaxHeap) Swap(i, j int) MaxHeap {
	a := heap[i]
	b := heap[j]

	a.Index = j
	b.Index = i
	heap[j] = a
	heap[i] = b
	return heap
}

// HasGreaterChild Returns true if there is one
func (heap MaxHeap) HasGreaterChild(ho HeapObject) bool {
	left := ho.Index * 2
	right := ho.Index*2 + 1
	if left < len(heap) {
		if c := heap[left]; c.Score > ho.Score {
			return true
		}
	}
	if right < len(heap) {
		if c := heap[right]; c.Score > ho.Score {
			return true
		}
	}
	return false
}

// GreaterThanParent Returns true if it is so
func (heap MaxHeap) GreaterThanParent(ho HeapObject) bool {
	if ho.Index == 1 {
		return false
	}
	p := heap[ho.Index/2]

	return ho.Score > p.Score
}

// GreaterChildIndex Returns the index of the greater child
func (heap MaxHeap) GreaterChildIndex(ho HeapObject) (int, bool) {
	if !heap.HasGreaterChild(ho) {
		return 0, false
	}
	left := ho.Index * 2
	right := ho.Index*2 + 1
	if left < len(heap) {
		if right < len(heap) && heap[right].Score > heap[left].Score {
			return right, true
		}
		// >, ==, or right is nil/out-of-range
		return left, true
	}
	return 0, false
}

// Pop Return the head and re-sorted heap
func (heap MaxHeap) Pop() (HeapObject, MaxHeap) {
	oldRoot := heap[ROOT]
	if len(heap) == 2 {
		return oldRoot, append(NewMaxHeap(), heap[len(heap)-1])
	}
	newRoot := heap[len(heap)-1]
	newRoot.Index = ROOT
	newHeap := NewMaxHeap()
	newHeap = append(newHeap, newRoot)
	newHeap = append(newHeap, heap[2:len(heap)-1]...)
	for newHeap.HasGreaterChild(newRoot) {
		greaterChildIndex, _ := newHeap.GreaterChildIndex(newRoot)
		newHeap = heap.Swap(greaterChildIndex, newRoot.Index)
		newRoot = heap[greaterChildIndex]
	}

	return oldRoot, newHeap
}

// Insert Add a new member
func (heap MaxHeap) Insert(ho HeapObject) MaxHeap {
	ho.Index = len(heap)
	newHeap := append(heap, ho)
	if ho.Index == 1 {
		return newHeap
	}

	for newHeap.GreaterThanParent(ho) {
		newHeap = newHeap.Swap(ho.Index, ho.Index/2)
		ho = newHeap[ho.Index/2]
	}
	return newHeap
}

func (heap MaxHeap) display() {
	fmt.Println("--HEAP STATE--")
	for i, ho := range heap {
		fmt.Printf("\t %d -> %d.Score[%d]\n", i, ho.Index, ho.Score)
	}
}

// NewMaxHeap Returns an empty heap
func NewMaxHeap() MaxHeap {
	return MaxHeap{
		HeapObject{},
	}
}
