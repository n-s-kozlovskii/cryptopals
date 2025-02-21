package main

import (
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
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

func main() {
	src := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

	sEnc, err := HexToBase64(src)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sEnc)
}
