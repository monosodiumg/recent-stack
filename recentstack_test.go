package recentstack_test

import (
	"monosodiumg/ds"
	"testing"
)

func TestNew(t *testing.T) {

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "err",
			wantErr: true,
		},
		
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := recentstack.New(); (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
