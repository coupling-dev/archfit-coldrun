package store

import "testing"

func TestSetRejectsEmptyKey(t *testing.T) {
	if err := New().Set("", "v"); err == nil {
		t.Fatal("Set with empty key: want error")
	}
}
