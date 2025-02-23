package set1

import "testing"

func TestDecypherWholeFile(t *testing.T) {
	want := "Now that the party is jumping"
	got := DecypherWholeFileSolver("../testdata/task4_input.txt")
	if got != want {
		t.Errorf("DecypherWholeFileSolver() = %q, want %q", got, want)
	}
}
