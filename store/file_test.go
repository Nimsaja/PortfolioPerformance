package store

import "testing"

func TestCreatePath(t *testing.T) {
	s := " Test Me "

	p := createPath(s)

	if p != "testme.txt" {
		t.Errorf("Expected path: %v, got %v", "testme", p)
	}
}
