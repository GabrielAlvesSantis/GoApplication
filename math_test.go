package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(5, 5)
	if total != 10 {
		t.Errorf("Soma(5, 5) = %d; want 10", total)
	}
}
