package set1

import (
	"fmt"
	"os"
)

func CountDifferentBits(a, b byte) int {
	// fmt.Println(strconv.FormatInt(int64(a), 2))
	// fmt.Println(strconv.FormatInt(int64(b), 2))
	var count int
	for i := 0; i < 8; i++ {
		a1 := a >> i
		a2 := a1 & 1

		b1 := b >> i
		b2 := b1 & 1
		res := a2 != b2
		if res {
			count += 1
		}
		// fmt.Printf(
		// 	"i=%d, a1=%s, a2=%s\n",
		// 	i,
		// 	strconv.FormatInt(int64(a1), 2),
		// 	strconv.FormatInt(int64(a2), 2),
		// )
		// fmt.Printf(
		// 	"i=%d, b1=%s, b2=%s\n",
		// 	i,
		// 	strconv.FormatInt(int64(b1), 2),
		// 	strconv.FormatInt(int64(b2), 2),
		// )
		// fmt.Printf("res:%t\n\n", res)
	}
	// fmt.Printf("count:%d\n", count)
	return count
}

func HammingDistance(a, b []byte) int {
	var count int
	for i, c := range a {
		count += CountDifferentBits(c, b[i])
		// fmt.Print(strconv.FormatInt(int64(c), 10))
		// fmt.Print(" ")
	}
	// fmt.Printf("final count: %d\n", count)
	return count
}

func EstimatedKeySizeAvg(a []byte, keySize int) float32 {
	//todo len a %  keySize == 0
	var acc int
	for i := keySize; i < len(a)-keySize; i += keySize {
		x := a[(i - keySize):i]
		y := a[i : i+keySize]
		distance := HammingDistance(x, y)
		acc += distance
		fmt.Printf("x: %s\ty: %s, d(x,y)=%d\n", x, y, distance)
	}
	res := float32(acc) / float32(2)
	fmt.Printf("EstimatedKeySize(%d)=%f\n", keySize, res)
	return res
}

func EstimatedKeySize(a []byte, keySize int) float32 {
	x := a[:keySize]
	y := a[keySize:]
	distance := HammingDistance(x, y)
	fmt.Printf("keySize: %d; x: %s\ty: %s, d(x,y)=%d\n", keySize, x, y, distance)
	return float32(distance) / float32(keySize)
}

func PrepareBytes(src []byte, keySize, numBlock int) []byte {
	//todo check len
	return src[:(keySize * numBlock)]
}

func FindSmallestKeySizes(src []byte) (int, int, int) {
	var s3, s2, s1 float32 = 100, 100, 100

	var i3, i2, i1 int

	for i := 2; i < 50; i++ {
		d := EstimatedKeySize(PrepareBytes(src, i, 2), i)
		//todo wft
		if d < s1 {
			s1 = s2
			s2 = s3
			s3 = d

			i1 = i2
			i2 = i3
			i3 = i
		} else if d < s2 {
			s2 = s3
			s3 = d
			i2 = i3
			i3 = i
		} else if d < s3 {
			s3 = d
			i3 = i
		}
	}
	fmt.Printf("d[%d]=%f; d[%d]=%f; d[%d]=%f\n", i1, s1, i2, s2, i3, s3)
	return i1, i2, i3
}

func SplitFile(filename string, blockSize int) [][]byte {
	dat, err := os.ReadFile(filename)
	check(err)

	m := 0
	// todo figure out what to do if file does not equaly separated
	// into blocks (works for blockSize = 7 for now)
	res := make([][]byte, len(dat)/blockSize)
	for j := blockSize; j <= len(dat); j += blockSize {
		res[m] = dat[(j - blockSize):(j)]
		m += 1
	}
	for i, block := range res {
		fmt.Printf("%d: %s\n", i, block)
	}

	return res
}

func Transponse(block [][]byte) [][]byte {
	//todo safeguards
	var newLen = len(block[0])
	res := make([][]byte, newLen)
	for i := 0; i < len(block); i++ {
		newRow := make([]byte, len(block))
		for j := 0; j < newLen; j++ {

		}
	}
}
