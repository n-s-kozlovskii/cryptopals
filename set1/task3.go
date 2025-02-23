package set1

import (
	"fmt"
	"log"
	"strings"
	"unicode"
)

func FillToLength(n int, b byte) []byte {
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = b
	}

	return res
}
func Decypher(input string) (string, rune) {
	alphabet := " !\"#$%&\\'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
	dst := LiteralToHex(input)
	n := len(dst)
	var bestDecypher string
	var bestScore int
	var key rune

	for i := 0; i < len(alphabet); i++ {
		data := FillToLength(n, alphabet[i])
		xor, err := FixedXor(dst[:n], data)
		if err != nil {
			fmt.Println("FixedXor(dst[:n], data) failed")
			log.Fatal(err)
		}
		score := Score(string(xor))
		if score > bestScore {
			bestDecypher = string(xor)
			bestScore = score
			key = rune(alphabet[i])
		}
	}
	return bestDecypher, key
}

func DecypherSolver(input string) string {
	bestDecypher, _ := Decypher(input)
	fmt.Printf("task3 message: %s\n", bestDecypher)
	return bestDecypher
}

func Score(a string) int {
	a = strings.ToUpper(a)
	//todo init map only once
	frequency := map[rune]int{
		'E': 28,
		'T': 27,
		'A': 26,
		'O': 25,
		'I': 24,
		'N': 23,
		'S': 22,
		'R': 21,
		'H': 20,
		'D': 17,
		'L': 16,
		'U': 15,
		'C': 14,
		'M': 13,
		'F': 12,
		'Y': 11,
		'W': 10,
		'G': 9,
		'P': 8,
		'B': 7,
		'V': 6,
		'K': 5,
		'X': 4,
		'Q': 3,
		'J': 2,
		'Z': 1,
		' ': 1,
	}
	var score int
	for _, ch := range a {
		score += frequency[unicode.ToUpper(ch)] //todo mb better extend frequency
	}

	return score
}
