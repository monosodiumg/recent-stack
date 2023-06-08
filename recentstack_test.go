package recentstack_test

import (
	"reflect"
	"testing"
	rs "monosodiumg/ds"
)

type (
	bb = []byte
)
func TestNew(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want rs.RecentStack[bb]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rs.New[bb](tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
