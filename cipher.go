package gutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"errors"
)

func Padding(source []byte) (dst []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(r.(string))
			dst = nil
		}
	}()
	diff := len(source) % aes.BlockSize
	paddingLength := aes.BlockSize - diff
	buffer := bytes.Buffer{}
	// this function is not gonna return an error
	// it would panic with something going wrong
	buffer.Write(source)
	buffer.Write(make([]byte, paddingLength-1))
	buffer.WriteByte(byte(int8(paddingLength)))
	dst = buffer.Bytes()
	return
}

func RemovePadding(source []byte) (dst []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(r.(string))
			dst = nil
		}
	}()
	totalLen := len(source)
	paddingLength := int(source[totalLen-1])
	dst = source[:totalLen-paddingLength]
	return
}

func HmacVerify(data, sign, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	if _, err := mac.Write(data); err != nil {
		return false
	}
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(expectedMAC, sign)
}

func HmacSign(data, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}

func CBCEncrypt(plaintext, iv, key []byte) (encrypted []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(r.(string))
			encrypted = nil
		}
	}()
	if block, err := aes.NewCipher(key); err == nil {
		if plaintext, err = Padding(plaintext); err == nil {
			blockMode := cipher.NewCBCEncrypter(block, iv)
			encrypted = make([]byte, len(plaintext))
			blockMode.CryptBlocks(encrypted, plaintext)
		}
	}
	return
}

func CBCDecrypt(encrypted, iv, key []byte) (plaintext []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(r.(string))
			encrypted = nil
		}
	}()
	if block, err := aes.NewCipher(key); err == nil {
		blockMode := cipher.NewCBCDecrypter(block, iv)
		plaintext = make([]byte, len(encrypted))
		blockMode.CryptBlocks(plaintext, encrypted)
		plaintext, err = RemovePadding(plaintext)
	}
	return
}
