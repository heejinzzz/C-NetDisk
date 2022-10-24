package handler

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var secretKey = "cNetDiskcNetDisk"

func AesEncrypt(origin string) string {
	// 转成字节数组
	data := []byte(origin)
	key := []byte(secretKey)

	// 分组秘钥
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error() + ". key 长度必须为 16/24/32")
	}

	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	data = PKCS7Padding(data, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	// 创建数组
	encrypted := make([]byte, len(data))
	// 加密
	blockMode.CryptBlocks(encrypted, data)
	//使用RawURLEncoding 不要使用StdEncoding
	//不要使用StdEncoding  放在url参数中回导致错误
	return base64.RawURLEncoding.EncodeToString(encrypted)
}

func AesDecrypt(encrypted string) string {
	//使用RawURLEncoding 不要使用StdEncoding
	//不要使用StdEncoding  放在url参数中回导致错误
	encryptedByte, _ := base64.RawURLEncoding.DecodeString(encrypted)
	k := []byte(secretKey)

	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		panic(err.Error() + ". key 长度必须为 16/24/32")
	}

	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	origin := make([]byte, len(encryptedByte))
	// 解密
	blockMode.CryptBlocks(origin, encryptedByte)
	// 去补全码
	origin = PKCS7UnPadding(origin)
	return string(origin)
}

// PKCS7Padding 补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

// PKCS7UnPadding 去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
