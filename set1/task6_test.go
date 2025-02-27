package set1

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	input1 := []byte("this is a test")
	input2 := []byte("wokka wokka!!!")
	got := HammingDistance(input1, input2)
	want := 37
	if got != want {
		t.Errorf("Distance() = %q, want %q", got, want)
	}
}

func TestEstimatedKeySize(t *testing.T) {
	b64Data, err := os.ReadFile("../testdata/task6_distance.txt")
	if err != nil {
		t.Errorf("TestHammingBase data corrupted, %q", err)
	}
	data := make([]byte, base64.StdEncoding.DecodedLen(len(b64Data)))
	base64.StdEncoding.Decode(data, b64Data)
	// fmt.Printf("as literal: %x\n", input1)
	fmt.Printf("from file: %x\n", data)

	got := EstimatedKeySize(data, 14)
	var want float32 = 37.0 / 14
	if got != want {
		t.Errorf("EstimatedKeySize()=%f, want %f", got, want)
	}
}

func TestPrepareBytes(t *testing.T) {
	b64Data, err := os.ReadFile("../testdata/task6_distance.txt")
	if err != nil {
		t.Errorf("TestHammingBase data corrupted, %q", err)
	}
	data := make([]byte, base64.StdEncoding.DecodedLen(len(b64Data)))
	base64.StdEncoding.Decode(data, b64Data)

	got := PrepareBytes(data, 14, 2)
	if string(got) != string(data[:28]) {
		t.Errorf("PrepareBytes()=%q, want %q", got, data[:28])
	}
}

func TestFindSmallestKeySizes(t *testing.T) {
	b64Data, err := os.ReadFile("../testdata/task6_input.txt")
	if err != nil {
		t.Errorf("TestHammingBase data corrupted, %q", err)
	}
	data := make([]byte, base64.StdEncoding.DecodedLen(len(b64Data)))
	base64.StdEncoding.Decode(data, b64Data)

	i1, i2, i3 := FindSmallestKeySizes(data)
	fmt.Printf("%d, %d, %d", i1, i2, i3)
}

func TestSplitFile(t *testing.T) {
	input := []byte{1, 1, 1, 2, 2, 2, 3, 3}
	want := [][]byte{
		{1, 1, 1},
		{2, 2, 2},
		{3, 3, 0},
	}
	got := SplitFile(input, 3)
	if !arraysEqual(got, want) {
		t.Errorf("SplitFile()=%q, want %q", got, want)
	}
}

func arraysEqual(this, that [][]byte) bool {
	if len(this) != len(that) {
		return false
	}

	for i, n := range this {
		if !bytes.Equal(n, that[i]) {
			return false
		}
	}
	return true
}

func TestTransponse(t *testing.T) {
	input := [][]byte{
		{0, 1, 2, 3}, /*  initializers for row indexed by 0 */
		{4, 5, 6, 7}, /*  initializers for row indexed by 1 */
		{8, 9, 10, 11},
	}

	want := [][]byte{
		{0, 4, 8},
		{1, 5, 9},
		{2, 6, 10},
		{3, 7, 11},
	}

	got := Transponse(input)

	if !arraysEqual(got, want) {
		t.Errorf("Transponse() = %q, want %q", got, want)
	}
}

func TestConstrunctKey(t *testing.T) {
	want := []byte("iiiinirnrn")
	datRaw, err := os.ReadFile("../testdata/task6_input.txt")
	if err != nil {
		t.Error("TestBlaBla readfile failed")
	}
	dat := make([]byte, base64.StdEncoding.DecodedLen(len(datRaw)))
	base64.StdEncoding.Decode(dat, datRaw)
	got := ConstrunctKey(dat, 10)
	if string(got) != string(want) {
		t.Errorf("BlaBla()=%q, want %q", got, want)
	}

}

func TestSuperF(t *testing.T) {
	SuperF("../testdata/task6_input.txt")
}
