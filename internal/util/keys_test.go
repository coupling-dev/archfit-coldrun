package util

import "testing"

func TestValidKey(t *testing.T) {
	tests := []struct {
		key  string
		want bool
	}{
		{"", false},
		{"a", true},
	}
	for _, tt := range tests {
		if got := ValidKey(tt.key); got != tt.want {
			t.Errorf("ValidKey(%q) = %v, want %v", tt.key, got, tt.want)
		}
	}
}
