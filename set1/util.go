package set1

import (
	"encoding/hex"
	"log"
)

func LiteralToHex(src string) []byte {
	return BytesToHex([]byte(src))
}

func BytesToHex(src []byte) []byte {
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}
	return dst[:n]
}
