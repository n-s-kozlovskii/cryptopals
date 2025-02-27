package set1

import (
	b64 "encoding/base64"
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
		fmt.Printf("HammingDistance(\"%x\", \"%x\")=%d", x, y, distance)
	}
	res := float32(acc) / float32(2)
	//fmt.Printf("EstimatedKeySize(%d)=%f\n", keySize, res)
	return res
}

func EstimatedKeySize(a []byte, keySize int) float32 {
	x := a[:keySize]
	y := a[keySize:]
	hD := HammingDistance(x, y)
	distance := float32(hD) / float32(keySize)
	return distance
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
		fmt.Printf("FindSmallestKeySizes i=%d, d=%f\n", i, d)
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
	// fmt.Printf("FindSmallestKeySizes d[%d]=%f; d[%d]=%f; d[%d]=%f\n", i1, s1, i2, s2, i3, s3)
	return i1, i2, i3
}

func SplitFile(dat []byte, blockSize int) [][]byte {
	fmt.Printf("SplitFile in %d len blocks, data len %d\n", blockSize, len(dat))
	m := 0
	res := make([][]byte, len(dat)/blockSize+1)

	for j := blockSize; j <= len(dat); j += blockSize {
		res[m] = dat[(j - blockSize):(j)]
		m += 1
	}
	res[m] = make([]byte, blockSize)
	if len(dat)%blockSize != 0 {
		for i, b := range dat[len(dat)/blockSize*blockSize:] {
			res[m][i] = b
		}
	}
	return res
}

func Transponse(block [][]byte) [][]byte {
	//todo safeguards
	N := len(block)
	M := len(block[0])
	res := make([][]byte, M)
	//todo make one go
	for i := 0; i < M; i++ {
		res[i] = make([]byte, N)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			res[j][i] = block[i][j]
		}
	}

	return res
}

func ConstrunctKey(dat []byte, blockSize int) []byte {
	fmt.Printf("ConstrunctKey %d block size, %d bytes\n", blockSize, len(dat))
	blocks := SplitFile(dat, blockSize)
	blocksT := Transponse(blocks)
	finalKey := make([]byte, blockSize)
	finalScore := 0
	for i, b := range blocksT {
		decripted, key, score := Decypher(b)
		fmt.Printf("Decypher %d) d: %q; key: %q\n", i, decripted, key)
		finalKey[i] = byte(key)
		finalScore += score
	}
	fmt.Printf("ConstrunctKey result: blockSize %d, score %f, key: %q\n", blockSize, float32(finalScore)/float32(blockSize), finalKey)
	return finalKey
}

func SuperF(filename string) {
	datRaw, err := os.ReadFile(filename)
	check(err)
	dat := make([]byte, b64.StdEncoding.DecodedLen(len(datRaw)))
	b64.StdEncoding.Decode(dat, datRaw)
	// FindSmallestKeySizes(datHex)

	keylen := 29
	key := ConstrunctKey(dat, keylen)
	fmt.Printf("1) %d\t%q\n", keylen, key)
	res := RepeatingKeyXor(dat, key)
	fmt.Printf("RepeatingKeyXor: %s\n", res)
}
