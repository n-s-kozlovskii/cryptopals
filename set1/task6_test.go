package set1

import "testing"

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
	input := []byte("HUIfTQsPAh9PE048")
	EstimatedKeySize(input, 4)
}

func TestFindSmallestKeySizes(t *testing.T) {
	input := []byte("HUIfTQsPAh9PE048GmllH0kcDk4TAQsHThsBFkU2AB4BSWQgVB0dQzNTTmVS" +
		"BgBHVBwNRU0HBAxTEjwMHghJGgkRTxRMIRpHKwAFHUdZEQQJAGQmB1MANxYG" +
		"DBoXQR0BUlQwXwAgEwoFR08SSAhFTmU+Fgk4RQYFCBpGB08fWXh+amI2DB0P" +
		"QQ1IBlUaGwAdQnQEHgFJGgkRAlJ6f0kASDoAGhNJGk9FSA8dDVMEOgFSGQEL" +
		"QRMGAEwxX1NiFQYHCQdUCxdBFBZJeTM1CxsBBQ9GB08dTnhOSCdSBAcMRVhI" +
		"CEEATyBUCHQLHRlJAgAOFlwAUjBpZR9JAgJUAAELB04CEFMBJhAVTQIHAh9P" +
		"G054MGk2UgoBCVQGBwlTTgIQUwg7EAYFSQ8PEE87ADpfRyscSWQzT1QCEFMa" +
		"TwUWEXQMBk0PAg4DQ1JMPU4ALwtJDQhOFw0VVB1PDhxFXigLTRkBEgcKVVN4")

	FindSmallestKeySizes(input)
}

func TestSplitFile(t *testing.T) {
	SplitFile("../testdata/task6_input.txt", 7)
}
