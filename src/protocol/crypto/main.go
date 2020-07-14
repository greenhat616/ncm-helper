package crypto

// Crypto is a transplanted library,
// from https://github.com/Binaryify/NeteaseCloudMusicApi/blob/master/util/crypto.js

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"io"
	"math/big"
)

var (
	iv          = []byte("0102030405060708")
	presetKey   = []byte("0CoJUm6Qyw8W8jud")
	linuxApiKey = []byte("rFgB&h#%2?^eDg:Q")
	base62      = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	publicKey   = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDgtQn2JZ34ZC28NWYpAUd98iZ37BUrX/aKzmFbt7clFSs6sXqHauqKWqdtLkF2KexO40H1YTX8z2lSgBBOAxLsvaklV8k4cBFK9snQXE9/DDaFt6Rr7iVZMldczhC0JNgTz+SHXT6CBHuX3e9SdB1Ua44oncaTWz7OBGLbCiK45wIDAQAB
-----END PUBLIC KEY-----`
	eapiKey = "e82ckenh8dichen8"
)
var util IUtil = &Util{}

func AESEncrypt(buffer []byte, mode string, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if mode == "gcm" {
		c, err := cipher.NewGCM(block)
		if err != nil {
			return nil, err
		}
		nonce := make([]byte, 12)
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			return nil, err
		}

		cipherText := c.Seal(nil, nonce, buffer, nil)
		return cipherText, nil
	} else if mode == "cbc" {
		buffer = util.PKCS7Padding(buffer, block.BlockSize())
		cipherText := make([]byte, len(buffer))
		c := cipher.NewCBCEncrypter(block, iv)
		c.CryptBlocks(cipherText, buffer)
		return cipherText, nil
	} else if mode == "ecb" {
		buffer = util.PKCS7Padding(buffer, block.BlockSize())
		encrypted := make([]byte, len(buffer))
		size := block.BlockSize()
		for bs, be := 0, size; bs < len(buffer); bs, be = bs+size, be+size {
			block.Encrypt(encrypted[bs:be], buffer[bs:be])
		}
		return encrypted, nil
	}
	return nil, errors.New("mismatch mode")
}

func RSAEncryptWithNoPadding(buffer []byte, key string) ([]byte, error) {
	// 解密 pem 格式的公钥
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)

	// 加密
	c := new(big.Int).SetBytes(buffer)
	return c.Exp(c, big.NewInt(int64(pub.E)), pub.N).Bytes(), nil

}

func Decrypt(buffer []byte) (decrypted []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknown error")
			}
		}
	}()
	cipher, _ := aes.NewCipher([]byte(eapiKey))
	decrypted = make([]byte, len(buffer))
	size := 16

	for bs, be := 0, size; bs < len(buffer); bs, be = bs+size, be+size {
		cipher.Decrypt(decrypted[bs:be], buffer[bs:be])
	}
	return
}

func WEAPI(data string) (params []byte, encSecKey []byte, err error) {
	secretKey, err := util.GenRandomBytes(16)
	// fmt.Printf("%v", secretKey)
	if err != nil {
		return
	}
	for k, v := range secretKey {
		secretKey[k] = byte(util.charCodeAt(util.base62Encode(int(v)), 0))
	}
	// fmt.Printf("%v", secretKey)
	presetData, err := AESEncrypt([]byte(data), "cbc", presetKey, iv)
	if err != nil {
		return
	}
	// fmt.Printf("%v", presetData)
	presetDataBase64 := make([]byte, base64.StdEncoding.EncodedLen(len(presetData)))
	base64.StdEncoding.Encode(presetDataBase64, presetData)
	secretData, err := AESEncrypt(presetDataBase64, "cbc", secretKey, iv)
	if err != nil {
		return
	}
	params = make([]byte, base64.StdEncoding.EncodedLen(len(secretData)))
	base64.StdEncoding.Encode(params, secretData)
	encSecKeyBytes, err := RSAEncryptWithNoPadding(util.reverse(secretKey), publicKey)
	encSecKey = make([]byte, hex.EncodedLen(len(encSecKeyBytes)))
	hex.Encode(encSecKey, encSecKeyBytes)
	return
}

func LinuxAPI(data string) (eParams []byte, err error) {
	encrypted, err := AESEncrypt([]byte(data), "ecb", linuxApiKey, nil)
	if err != nil {
		return
	}
	eParams = make([]byte, hex.EncodedLen(len(encrypted)))
	hex.Encode(eParams, encrypted)
	eParams = bytes.ToUpper(eParams)
	return
}

func EAPI(url string, data string) (params []byte, err error) {
	msg := "nobody" + url + "use" + data + "md5forencrypt"
	h := md5.New()
	h.Write([]byte(msg))
	digest := hex.EncodeToString(h.Sum(nil))
	target := url + "-36cd479b6b5-" + data + "-36cd479b6b5-" + digest
	encrypted, err := AESEncrypt([]byte(target), "ecb", []byte(eapiKey), nil)
	if err != nil {
		return
	}
	params = make([]byte, hex.EncodedLen(len(encrypted)))
	hex.Encode(params, encrypted)
	params = bytes.ToUpper(params)
	return
}
