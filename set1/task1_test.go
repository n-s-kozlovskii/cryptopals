package set1

import "testing"

func TestHex2Base64(t *testing.T) {
	src := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	if got, _ := HexToBase64(src); got != want {
		t.Errorf("HexToBase64() = %q, want %q", got, want)
	}
}
