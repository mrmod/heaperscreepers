package minheap

import "testing"

func TestInsert(t *testing.T) {
	h := NewMinHeap()
	if len(h) != 1 {
		t.Fatal("Expected only one element (nil root)")
	}
	h = h.Insert(HeapObject{Score: 14})
	if o := h.Peek(); o.Score != 14 {
		t.Fatalf("Expected 14, got %d", o.Score)
	}
	h = h.Insert(HeapObject{Score: 23})
	if o := h.Peek(); o.Score != 14 {
		t.Fatalf("Expected 14, got %d", o.Score)
	}
	h = h.Insert(HeapObject{Score: 9})
	if o := h.Peek(); o.Score != 9 {
		t.Fatalf("Expected 9, got %d, %v", o.Score, h)
	}
}
