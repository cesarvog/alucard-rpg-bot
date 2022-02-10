package main 

import (
	"testing"
)

func TestRoll(t *testing.T) {
	d := &Dice{}

	s := d.Roll(5, 4);

	if s == "" {
		t.Error("No results")
	}
}

func BenchmarkRoll(b *testing.B) {
	d := &Dice{}
	for i := 0; i < b.N; i++ {
		d.Roll(10, 10)
	}
}

