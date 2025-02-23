package set1

import (
	b64 "encoding/base64"
	"fmt"
)

func HexToBase64(src string) (string, error) {
	dst := LiteralToHex(src)
	fmt.Printf("task1 message: %s\n", dst)

	sEnc := b64.StdEncoding.EncodeToString(dst)
	return sEnc, nil
}
