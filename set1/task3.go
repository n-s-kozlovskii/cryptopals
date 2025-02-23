package set1

import (
	"fmt"
	"log"
)

func FixedXor2(left, right []byte) (string, error) {
	n := len(left)

	res := make([]byte, n)

	for i := 0; i < n; i++ {
		res[i] = left[i] ^ right[i]
	}

	return string(res[:n]), nil
}

func FillToLength(n int, b byte) []byte {
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = b
	}

	return res
}
func Decypher(input string) []string {
	alphabet := "ETAONRISHDLFCMUGYPWBVKJXZQ"
	dst := LiteralToHex(input)
	n := len(dst)
	res := make([]string, len(alphabet))

	for i := 0; i < len(alphabet); i++ {
		data := FillToLength(n, alphabet[i])
		xor, err := FixedXor2(dst[:n], data)
		if err != nil {
			fmt.Println("FixedXor(dst[:n], data) failed")
			log.Fatal(err)
		}
		res[i] = string(xor)
	}
	fmt.Printf("task3 message: %s\n", res[23])
	return res
}
