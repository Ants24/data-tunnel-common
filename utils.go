package datatunnelcommon

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"errors"
	"os"
	"unicode"
)

func HasLowercase(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return true
		}
	}
	return false
}
func HasUppercase(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}

// 判断文件是否存在
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// 读取文件内容
func ReadFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// 加密
func Encrypt(text string, key string) (string, error) {
	src := []byte(text)
	keys := []byte(key)
	block, err := des.NewCipher(keys)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = ZeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

// 解密
func Decrypt(decrypted string, key string) (string, error) {
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return "", err
	}
	keys := []byte(key)
	block, err := des.NewCipher(keys)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)
	return string(out), nil
}
func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}
