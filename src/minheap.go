package heap

import "fmt"

// MinHeap A collection of HeapObject
type MinHeap []HeapObject

// LessThanParent True if ho.Index/2 Score less than our score
func (heap MinHeap) LessThanParent(ho HeapObject) bool {
	if ho.Index == 1 {
		return false
	}

	if parent := heap[ho.Index/2]; ho.Score < parent.Score {
		return true
	}
	return false
}

// LessThanIndex Returns true if the HeapObject at Index has a lower score than score
func (heap MinHeap) LessThanIndex(score, index int) bool {
	if index < len(heap) {
		if h := heap[index]; score < h.Score {
			return true
		}
	}
	return false
}

// HasLesserChild Returns true if either is less than HO
func (heap MinHeap) HasLesserChild(ho HeapObject) bool {
	leftIndex := ho.Index * 2
	rightIndex := ho.Index*2 + 1

	if leftIndex < len(heap) {
		if child := heap[leftIndex]; child.Score < ho.Score {
			return true
		}
	}
	if rightIndex < len(heap) {
		if child := heap[rightIndex]; child.Score < ho.Score {
			return true
		}
	}

	return false
}

// LesserChildIndex Returns the lesser child index and boolean if one exists at all
func (heap MinHeap) LesserChildIndex(ho HeapObject) (int, bool) {
	if !heap.HasLesserChild(ho) {
		return 0, false
	}
	var leftChild, rightChild *HeapObject
	leftIndex := ho.Index * 2
	rightIndex := ho.Index*2 + 1
	if leftIndex < len(heap) {
		leftChild = &heap[leftIndex]
	}
	if rightIndex < len(heap) {
		rightChild = &heap[rightIndex]
	}

	if leftChild != nil {
		if rightChild != nil && rightChild.Score < leftChild.Score {
			return rightIndex, true
		}
		return leftIndex, true
	}
	return 0, false
}

// Swap Swap a into b's position and b to a's
func (heap MinHeap) Swap(i, j int) MinHeap {

	a := heap[i]
	b := heap[j]
	a.Index = j
	b.Index = i
	heap[i] = b
	heap[j] = a

	return heap
}
func (heap MinHeap) display() {
	fmt.Println("--HEAP STATE--")
	for i, ho := range heap {
		fmt.Printf("\t %d -> %d.Score[%d]\n", i, ho.Index, ho.Score)
	}
}

// Insert Add a new member
func (heap MinHeap) Insert(ho HeapObject) MinHeap {
	ho.Index = len(heap)
	newHeap := append(heap, ho)
	if ho.Index == 1 {
		return newHeap
	}

	for newHeap.LessThanParent(ho) {
		newHeap = newHeap.Swap(ho.Index, ho.Index/2)
		ho = newHeap[ho.Index/2]
	}
	return newHeap
}

// Pop Return the Root
func (heap MinHeap) Pop() (HeapObject, MinHeap) {
	oldRoot := heap[1]
	newHeap := NewMinHeap()

	newRoot := heap[len(heap)-1]
	newRoot.Index = 1
	// Root the tree at the newRoot
	newHeap = append(newHeap, newRoot)
	// The deepest element should be ommitted from the new slice
	newHeap = append(newHeap, heap[2:len(heap)-1]...)
	// While the moved element Score is > than at least one child
	for newHeap.HasLesserChild(newRoot) {
		childIndex, _ := newHeap.LesserChildIndex(newRoot)
		newHeap = newHeap.Swap(childIndex, newRoot.Index)
		newRoot = newHeap[childIndex]
	}
	return oldRoot, newHeap
}

// Peek Show the root
func (heap MinHeap) Peek() HeapObject {
	return heap[ROOT]
}

// NewMinHeap Min is always at the head
func NewMinHeap() MinHeap {
	return MinHeap{HeapObject{}}
}
