package set1

import (
	"bytes"
	"testing"
)

func TestFixedXor(t *testing.T) {
	want := []byte("746865206b696420646f6e277420706c6179")

	if got, _ := FixedXor(
		[]byte("1c0111001f010100061a024b53535009181c"),
		[]byte("686974207468652062756c6c277320657965"),
	); !bytes.Equal(got, want[:]) {
		t.Errorf("FixedXor() = %q, want %q", got, want)
	}
}

func TestFixedXor_error(t *testing.T) {
	if _, err := FixedXor(
		[]byte("1c0111001f010100061a024b5353500918"),
		[]byte("686974207468652062756c6c277320657965"),
	); err == nil {
		t.Errorf("FixedXor() must fail on arrays with arrays different length, but err is nil")
	}
}
