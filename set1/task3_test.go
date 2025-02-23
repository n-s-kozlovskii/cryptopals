package set1

import (
	"testing"
)

func TestDecypher(t *testing.T) {
	want := "Cooking MC's like a pound of bacon"
	got := Decypher("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	found := false
	for _, decyphered := range got {
		found = found || decyphered == want
	}
	if !found {
		t.Errorf("Decypher() = %q, want %q", got, want)
	}
}
