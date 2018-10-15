package heap

// HeapObject Stores a something
type HeapObject struct {
	Index, Score int
	Value        interface{}
}

// ROOT is the root node
const ROOT = 1
