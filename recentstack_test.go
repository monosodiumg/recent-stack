package recentstack_test

import (
	rs "monosodiumg/ds"
	"testing"
)

type (
	bb = []byte
)

func TestNew(t *testing.T) {
	// type args struct {
	// 	capacity int
	// }
	// tests := []struct {
	// 	name string
	// 	args args
	// 	want rs.RecentStack[bb]
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := rs.New[bb](tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("New() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }

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
		if err != nil || s1 == nil   {
			t.Errorf("New(2) = (%v, %e), want = (not nil, nil)", s1, err)
		} else {
		if s1.Length() != 0 {
			t.Errorf("New(2).Length want 0, got %d", s1.Length() )
		}
	}
	})

	// fail if capacity does not match arg
	s1, err = rs.New[bb](2)
	bb1,bb2,bb3 := []byte("a"), []byte("b"), []byte("c") 
	s1.Push(bb1)
	l1 := s1.Length()
	s1.Push(bb2)
	l2 := s1.Length()
	s1.Push(bb3)
	l3 := s1.Length()

	t.Run("New(2)", func(t *testing.T) {
		if l1!= 1 || l2 != 2 || l3 != 2 {
			t.Errorf("New(2) with three pushes expected lengths of 1,2,2, got %d,%d,%d", l1,l2,l3)
		}
	})

}
