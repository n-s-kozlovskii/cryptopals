package set1

import (
	"encoding/hex"
	"log"
)

func LiteralToHex(a string) []byte {
	src := []byte(a)
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}
	return dst[:n]
}
