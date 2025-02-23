package set1

import (
	"encoding/hex"
	"fmt"
)

func FixedXor(left, right []byte) ([]byte, error) {
	n := len(left)
	if n != len(right) {
		//todo what to return on left?
		return []byte{}, fmt.Errorf("arrays lenght missmatched, %d != %d", len(left), len(right))
	}

	res := make([]byte, n)

	for i := 0; i < n; i++ {
		res[i] = left[i] ^ right[i]
	}

	return res[:n], nil

}

func FixedHexXor(left, right []byte) ([]byte, error) {
	res, err := FixedXor(BytesToHex(left), BytesToHex(right))
	fmt.Printf("task2 message: %s\n", res)
	if err != nil {
		return []byte{}, err
	}
	res2 := make([]byte, hex.EncodedLen(len(res)))
	hex.Encode(res2, res)
	return res2, nil
}
