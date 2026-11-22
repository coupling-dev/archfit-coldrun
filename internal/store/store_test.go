package store

import "testing"

func TestSetRejectsEmptyKey(t *testing.T) {
	if err := New().Set("", "v"); err == nil {
		t.Fatal("Set with empty key: want error")
	}
}

func TestDeleteRemovesKey(t *testing.T) {
	s := New()
	_ = s.Set("k", "v")
	s.Delete("k")
	if _, ok := s.Lookup("k"); ok {
		t.Fatal("Delete(k): key still present")
	}
}
