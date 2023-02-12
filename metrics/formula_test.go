package metrics

import (
	"testing"
)

func TestValidNetScore(t *testing.T) {
	// test the compute net score function with a
	// factor that should make net score 0
	f1 := Factor{
		Weight:       1,
		Value:        1,
		AllOrNothing: false,
	}
	f2 := Factor{
		Weight:       1,
		Value:        0,
		AllOrNothing: true,
	}
	fs := []Factor{f1, f2}
	res := ComputeNetScore(fs)

	if res != 0 {
		t.Errorf("res should be 0")
	}
}

func TestManagerBadInput(t *testing.T) {
	// test the capping of bad factor values [0,1]
	f1 := Factor{
		Weight:       45,
		Value:        4724,
		AllOrNothing: false,
	}
	f2 := Factor{
		Weight:       -1,
		Value:        999,
		AllOrNothing: false,
	}
	fs := []Factor{f1, f2}
	res := ComputeNetScore(fs)

	if res < 0 || res > 1 {
		t.Errorf("res should be capped to 0 & 1")
	}
}
