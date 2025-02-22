package set1

import (
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
)

func HexToBase64(src []byte) (string, error) {
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		return "", err
	}
	fmt.Printf("%s\n", dst[:n])

	sEnc := b64.StdEncoding.EncodeToString(dst[:n])
	return sEnc, nil
}
