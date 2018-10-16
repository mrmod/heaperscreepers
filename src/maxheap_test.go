package heap

import "testing"

func TestMaxInsert(t *testing.T) {
	h := NewMaxHeap()
	h = h.Insert(HeapObject{Score: 30})
	h = h.Insert(HeapObject{Score: 25})
	if o := h.Peek(); o.Score != 30 {
		t.Fatalf("Expected 30, got %d", o.Score)
	}
	h = h.Insert(HeapObject{Score: 40})
	h = h.Insert(HeapObject{Score: 10})
	if o := h.Peek(); o.Score != 40 {
		t.Fatalf("Expected 40, got %d", o.Score)
	}
}

func TestMaxPop(t *testing.T) {
	h := NewMaxHeap()
	h = h.Insert(HeapObject{Score: 30})
	h = h.Insert(HeapObject{Score: 25})
	h = h.Insert(HeapObject{Score: 40})
	h = h.Insert(HeapObject{Score: 10})
	ho, h := h.Pop()
	if ho.Score != 40 {
		t.Fatalf("Expected 40, got %d", ho.Score)
	}
	if o := h.Peek(); o.Score != 30 {
		t.Fatalf("Expected 30, got %d", o.Score)
	}
	if len(h) != 5 {
		t.Fatalf("Expected 5 items, found %d", len(h))
	}
}
