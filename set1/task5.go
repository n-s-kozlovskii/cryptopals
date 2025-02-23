package set1

func RepeatingKeyXor(textToEncrypt, key []byte) []byte {
	key_len := len(key)
	encrypted := make([]byte, len(textToEncrypt))
	for i, b := range textToEncrypt {
		encrypted[i] = b ^ key[i%key_len]
	}
	return encrypted
}
