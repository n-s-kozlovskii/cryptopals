package set1

import (
	"testing"
)

func TestDecypher(t *testing.T) {
	want := "Cooking MC's like a pound of bacon"
	got := DecypherSolver("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if got != want {
		t.Errorf("Decypher() = %q, want %q", got, want)
	}
}
