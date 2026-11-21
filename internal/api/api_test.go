package api

import "testing"

func TestPutThenGet(t *testing.T) {
	srv := New()
	if err := srv.Put("k", "v"); err != nil {
		t.Fatal(err)
	}
	got, ok := srv.Get("k")
	if !ok || got != "v" {
		t.Fatalf("Get(k) = %q, %v; want v, true", got, ok)
	}
}
