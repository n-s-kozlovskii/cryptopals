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

	left_dec := make([]byte, hex.DecodedLen(n))
	decoded_len, err := hex.Decode(left_dec, left)
	if err != nil {
		return []byte{}, err
	}

	right_dec := make([]byte, hex.DecodedLen(n))
	_, err1 := hex.Decode(right_dec, right)
	if err1 != nil {
		return []byte{}, err
	}
	res := make([]byte, decoded_len)

	for i := 0; i < decoded_len; i++ {
		res[i] = left_dec[i] ^ right_dec[i]
	}
	resHex := make([]byte, hex.EncodedLen(decoded_len))
	hex.Encode(resHex, res)

	fmt.Printf("task2 message: %s\n", res[:decoded_len])
	return resHex, nil

}
