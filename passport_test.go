package passport

import (
	"testing"
)

func BenchmarkRun(b *testing.B) {
	r := PreRun()
	for n := 0; n < b.N; n++ {
		r.Next()
	}
}

func TestRun(t *testing.T) {
	r := PreRun()
	current := r.Next()
	for n := 0; n < 10; n++ {
		new := r.Next()
		if new <= current {
			t.Errorf("New ID is equal or less than current id, current: %d, new : %d, iteration : %d ", current, new, n)
		}
		current = new
	}
}
