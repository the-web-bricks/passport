package passport

import (
	"testing"
)

func BenchmarkRun(b *testing.B) {
	passport := PreRun(defaultNodeIDs)
	for n := 0; n < b.N; n++ {
		r.ID()
	}
}

func TestRun(t *testing.T) {
	p := PreRun(defaultNodeID)
	current := p.ID()
	for n := 0; n < 10; n++ {
		new := p.ID()
		if new <= current {
			t.Errorf("New ID is equal or less than current id, current: %d, new : %d, iteration : %d ", current, new, n)
		}
		current = new
	}
}
