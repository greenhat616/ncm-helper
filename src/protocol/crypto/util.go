package crypto

import (
	"bytes"
	"math"
	"strconv"
	"strings"
)

func base62Encode(num int) string {
	baseStr := ""
	for {
		if num <= 0 {
			break
		}

		i := num % 62
		baseStr += base62[i]
		num = (num - i) / 62
	}
	return baseStr
}

func base62Decode(base string) int {
	rs := 0
	length := len(base)
	f := flip(base62)
	for i := 0; i < length; i++ {
		rs += f[string(base[i])] * int(math.Pow(62, float64(i)))
	}
	return rs
}

func flip(s []string) map[string]int {
	f := make(map[string]int)
	for index, value := range s {
		f[value] = index
	}
	return f
}

func charCodeAt(s string, n int) rune {
	i := 0
	for _, r := range s {
		if i == n {
			return r
		}
		i++
	}
	return 0
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

func genRandomBytes(size int) (blk []byte, err error) {
	blk = make([]byte, size)
	_, err = rand.Read(blk)
	return
}

func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

func reverse(s []byte) []byte {
	a := make([]byte, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

func convert(b []byte) string {
	s := make([]string, len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s, ",")
}
