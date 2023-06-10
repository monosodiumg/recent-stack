package recentstack

import (
	"errors"
	"sync"
)

type (
	// RecentStack is a limited-capacity stack that forgets the bottom
	// item when an item is pushed and the stack is full
	RecentStack[T any] interface {
		Push(T)
		Pop() (T, error)
		Length() int
	}

	recentStack[T any] struct {
		sync.RWMutex
		items []T
		// physical index where the first item in the stack is to be found or placed.
		bottom int
		// number of items on the stack
		length   int
		capacity int
	}
)

// ErrNoRows is returned by New when capacity is less than 2.
// For a capacity of 1 you might as well just use a variable.
var ErrTooSmall = errors.New("recentstack: minimum sensible capacity is 2")

// New returns an empty stack
func New[T any](capacity int) (RecentStack[T], error) {
	if capacity < 2 {
		return nil, ErrTooSmall
	}
	return &recentStack[T]{
		items:    make([]T, capacity),
		capacity: capacity,
	}, nil
}

// Length returns the number of items on the stack
func (r *recentStack[T]) Length() int {
	r.Lock()
	defer r.Unlock()
	// physical top: (bottom+top)
	return r.length
}

// Pushes an item onto the stack. If the stack is already full then the bottom item will be lost
func (r *recentStack[T]) Push(x T) {
	r.Lock()
	defer r.Unlock()
	r.items[(r.bottom+r.length)%r.capacity] = x
	if r.length == r.capacity {
		r.bottom = (r.bottom + 1) % r.capacity
	} else {
		r.length = r.length + 1
	}
}

var ErrEmpty = errors.New("Illegal pop - stack is empty")

func (r *recentStack[T]) Pop() (T, error) {
	var x T
	r.Lock()
	defer r.Unlock()
	if r.length == 0 {
		return x, ErrEmpty
	} else {
		x = r.items[(r.bottom+r.length-1)%r.capacity]
		r.length = (r.length - 1) % r.capacity
		return x, nil
	}
}
