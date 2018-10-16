package heap

import "testing"

func TestMinInsert(t *testing.T) {
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

	h = h.Insert(HeapObject{Score: 2})
	if o := h.Peek(); o.Score != 2 {
		t.Fatalf("Expected 2, got %d, %v", o.Score, h)
	}
}

func TestMinPop(t *testing.T) {
	h := NewMinHeap()
	h = h.Insert(HeapObject{Score: 14})
	h = h.Insert(HeapObject{Score: 23})
	h = h.Insert(HeapObject{Score: 9})
	h = h.Insert(HeapObject{Score: 90})
	h = h.Insert(HeapObject{Score: 17})
	h = h.Insert(HeapObject{Score: 2})
	if p := h.Peek(); p.Score != 2 {
		t.Fatalf("Expected to find 2, found %d", p.Score)
	}
	popped, h := h.Pop()
	if popped.Score != 2 {
		t.Fatal("Expected 2, got", popped.Score)
	}
	if len(h) != 6 {
		t.Fatalf("Expected only 6 elements, got %d, %#v", len(h), h)
	}
	if h[1].Score != 9 {
		t.Fatal("Expected the newRoot to be the smallest node, got", h[1].Score)
	}
}
