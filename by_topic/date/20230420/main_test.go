package main

import "testing"

func TestProbBoardedPosition(t *testing.T) {
	got := ProbBoardedPosition(0, 0, 1)
	want := 0.25
	if got != want {
		t.Errorf("ProbBoardedPosition(%d, %d, %d)=%g -- should be %g\n",
			0, 0, 1, got, want)
	}
}
