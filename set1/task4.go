package set1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func DecypherWholeFile(fileName string) string {
	f, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(f)
	var maxScore int
	var res string
	for scanner.Scan() {
		tmp, _ := DecypherString(scanner.Text())
		score := Score(tmp)
		if score > maxScore {
			maxScore = score
			res = tmp
		}
	}
	return strings.TrimSpace(res)
}

func DecypherWholeFileSolver(fileName string) string {
	res := DecypherWholeFile(fileName)
	fmt.Printf("task4 message: %s\n", res)
	return res
}
