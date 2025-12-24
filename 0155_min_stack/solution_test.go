package minstack

import "testing"

func TestMinStack(t *testing.T) {
	ms := Constructor()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)

	if ms.GetMin() != -3 {
		t.Errorf("GetMin() = %d, want -3", ms.GetMin())
	}

	ms.Pop()
	if ms.Top() != 0 {
		t.Errorf("Top() = %d, want 0", ms.Top())
	}

	if ms.GetMin() != -2 {
		t.Errorf("GetMin() = %d, want -2", ms.GetMin())
	}
}
