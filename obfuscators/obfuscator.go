package obfuscators

func Obfuscate(str string, salt []byte) []byte {
	result := make([]byte, 0, len(str))
	for i, char := range []byte(str) {
			result = append(result, char^salt[i%len(salt)])
	}
	return result
}

func Deobfuscate(bytes []byte, salt []byte) (string, error) {
	result := make([]byte, 0, len(bytes))
	for i, b := range bytes {
			result = append(result, b^salt[i%len(salt)])
	}
	return string(result), nil
}