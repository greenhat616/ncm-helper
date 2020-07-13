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
)

var (
	iv          = []byte("0102030405060708")
	presetKey   = []byte("0CoJUm6Qyw8W8jud")
	linuxApiKey = []byte("rFgB&h#%2?^eDg:Q")
	base62      = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	publicKey   = "-----BEGIN PUBLIC KEY-----\\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDgtQn2JZ34ZC28NWYpAUd98iZ37BUrX/aKzmFbt7clFSs6sXqHauqKWqdtLkF2KexO40H1YTX8z2lSgBBOAxLsvaklV8k4cBFK9snQXE9/DDaFt6Rr7iVZMldczhC0JNgTz+SHXT6CBHuX3e9SdB1Ua44oncaTWz7OBGLbCiK45wIDAQAB\\n-----END PUBLIC KEY-----"
	eapiKey     = "e82ckenh8dichen8"
)

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
		if len(buffer)%aes.BlockSize != 0 {
			return nil, errors.New("plaintext is not a multiple of the block size")
		}
		buffer = PKCS7Padding(buffer, block.BlockSize())
		c := cipher.NewCBCEncrypter(block, iv)
		cipherText := make([]byte, len(buffer))
		c.CryptBlocks(cipherText, buffer)
	} else if mode == "ecb" {
		buffer = PKCS7Padding(buffer, block.BlockSize())
		encrypted := make([]byte, len(buffer))
		size := block.BlockSize()
		for bs, be := 0, size; bs < len(buffer); bs, be = bs+size, be+size {
			block.Encrypt(encrypted[bs:be], buffer[bs:be])
		}
		return encrypted, nil
	}
	return nil, errors.New("mismatch mode")
}

func RSAEncrypt(buffer []byte, key string) ([]byte, error) {
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
	return rsa.EncryptPKCS1v15(rand.Reader, pub, buffer)
}

func Decrypt(buffer []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(generateKey([]byte(eapiKey)))
	if err != nil {
		return nil, err
	}
	decrypted := make([]byte, len(buffer))
	for bs, be := 0, cipher.BlockSize(); bs < len(buffer); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], buffer[bs:be])
	}
	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	return decrypted[:trim], nil
}

func WEAPI(data string) (params []byte, encSecKey []byte, err error) {
	secretKey, err := genRandomBytes(16)
	if err != nil {
		return
	}
	for k, v := range secretKey {
		secretKey[k] = byte(charCodeAt(base62Encode(int(v)), 0))
	}
	presetData, err := AESEncrypt([]byte(data), "cbc", presetKey, iv)
	if err != nil {
		return
	}
	presetDataBase64 := make([]byte, len(presetData))
	base64.StdEncoding.Encode(presetDataBase64, presetData)
	secretData, err := AESEncrypt(presetDataBase64, "cbc", secretKey, iv)
	if err != nil {
		return
	}
	base64.StdEncoding.Encode(params, secretData)
	encSecKey, err = RSAEncrypt(reverse(secretKey), publicKey)
	return
}

func LinuxAPI(data string) (eParams []byte, err error) {
	encrypted, err := AESEncrypt([]byte(data), "ecb", linuxApiKey, nil)
	if err != nil {
		return
	}
	hex.Encode(eParams, encrypted)
	bytes.ToUpper(eParams)
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
	hex.Encode(params, encrypted)
	bytes.ToUpper(params)
	return
}
