package set1

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"os"
)

//openssl enc -aes-128-ecb -d -a -in testdata/task7.txt -K "59454C4C4F57205355424D4152494E45"

func Aes128ebc() {
	datRaw, err := os.ReadFile("../testdata/task7.txt")
	check(err)
	Aes128ebc2(string(datRaw))
}

func Aes128ebc2(b64input string) {
	key := []byte("YELLOW SUBMARINE")
	keyLen := len(key)
	aes128, err := aes.NewCipher(key)
	check(err)
	coded := make([]byte, base64.RawStdEncoding.DecodedLen(len(b64input)))
	base64.StdEncoding.Decode(coded, []byte(b64input))
	for i := 0; i < len(coded)-keyLen; i += keyLen {
		blockToDecode := coded[i : i+keyLen]
		decrypt := make([]byte, keyLen)
		aes128.Decrypt(decrypt, blockToDecode)
		fmt.Printf("%s\n", string(decrypt))
	}
	// decrypt := make([]byte, keyLen)
	// aes128.Decrypt(decrypt, coded)
	// fmt.Printf("%s\n", string(decrypt))
	// fmt.Printf("%d", len(decrypt))
}
