package recentstack_test

import (
	rs "monosodiumg/ds"
	"reflect"
	"testing"
)

type (
	bb = []byte
)

func TestNew(t *testing.T) {
	// fail on unnacceptable capacity
	s1, err := rs.New[bb](1)
	t.Run("New(1)", func(t *testing.T) {
		if err == nil || s1 != nil {
			t.Errorf("New(1) = (%v, %e), want = (%v, nil)", s1, err, rs.ErrTooSmall)
		}
	})

	// fail if wrong response for acceptable capcity
	s1, err = rs.New[bb](2)
	t.Run("New(2)", func(t *testing.T) {
		if err != nil || s1 == nil {
			t.Errorf("New(2) = (%v, %e), want = (not nil, nil)", s1, err)
		} else {
			if s1.Length() != 0 {
				t.Errorf("New(2).Length want 0, got %d", s1.Length())
			}
		}
	})
}

func TestLength(t *testing.T) {
	// fail if capacity does not match arg
	t.Run("New(2)", func(t *testing.T) {
		s, _ := rs.New[bb](2)
		bb1, bb2, bb3 := []byte("a"), []byte("b"), []byte("c")
		s.Push(bb1)
		l1 := s.Length()
		s.Push(bb2)
		l2 := s.Length()
		s.Push(bb3)
		l3 := s.Length()
		if l1 != 1 || l2 != 2 || l3 != 2 {
			t.Errorf("New(2) with three pushes expected lengths of 1,2,2, got %d,%d,%d", l1, l2, l3)
		}
	})
}

func TestPop(t *testing.T) {
	// fail Pop does not return last pushed item, no wraparound
	t.Run("Pop() no wraparound", func(t *testing.T) {
		s, _ := rs.New[bb](2)
		bb1 := []byte("a")
		s.Push(bb1)
		x, e := s.Pop()
		if e != nil || !reflect.DeepEqual(bb1, x) {
			t.Errorf("Pop() want (%v,%v), got (%v,%v)", bb1, nil, x, e)
		}
	})

	// fail Pop does not return last pushed item, with wraparound
	t.Run("Pop() after wraparound", func(t *testing.T) {
		s, _ := rs.New[bb](2)
		bb1 := []byte("a")
		s.Push(bb1)
		bb2 := []byte("b")
		s.Push(bb2)
		bb3 := []byte("c")
		s.Push(bb3)
		x, e := s.Pop()
		if e != nil || !reflect.DeepEqual(bb3, x) {
			t.Errorf("Pop() want (%v,%v), got (%v,%v)", bb1, nil, x, e)
		}
	})

	// fail Pop on empty stack
	t.Run("Pop() on empty unused stack", func(t *testing.T) {
		s, _ := rs.New[bb](2)
		x, e := s.Pop()
		if e == nil || x != nil {
			t.Errorf("Pop() want (%v,%v), got (%v,%v)", nil, rs.ErrEmpty, x, nil)
		}
	})

	t.Run("Pop() on empty used stack", func(t *testing.T) {
		s, _ := rs.New[bb](2)
		bb1 := []byte("a")
		s.Push(bb1)
		s.Pop()
		x, e := s.Pop()
		if e == nil || x != nil {
			t.Errorf("Pop() want (%v,%v), got (%v,%v)", nil, rs.ErrEmpty, x, nil)
		}

	})

}
