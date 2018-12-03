package gutil

import (
	"crypto/aes"
	"testing"
)

func TestHMacSign(t *testing.T) {
	key := []byte("chrainlovescarol")
	data := []byte("carolloveschrain")
	sign := HmacSign(data, key)
	t.Logf(string(Base64Encode(sign)))
}

func TestHMacVerify(t *testing.T) {
	key := []byte("chrainlovescarol")
	data := []byte("carolloveschrain")
	sign := HmacSign(data, key)
	if !HmacVerify(data, sign, key) {
		panic("Verification failed")
	}
}

func TestPadding(t *testing.T) {
	data := []byte("carolloveschrain")
	if paddedData, err := Padding(data); err != nil {
		panic("Padding failed")
	} else {
		dataLen := len(data)
		paddingLen := aes.BlockSize - dataLen%aes.BlockSize
		if paddingLen != int(paddedData[len(paddedData)-1]) {
			panic("doesn't match")
		}
	}
}

func TestCBC(t *testing.T) {
	var data []byte
	var err error
	key := Md5Encode([]byte("chrainlovescarol"))
	data = []byte("carolloveschrain")
	iv := Md5Encode([]byte("123456"))
	if data, err = Padding(data); err != nil {
		t.Error(err.Error())
		panic("Padding string error")
	}
	var encrypted []byte
	if encrypted, err = CBCEncrypt(data, iv, key); err != nil {
		t.Error(err.Error())
		panic("Encrytion failed")
	}
	var plaintext []byte
	if plaintext, err = CBCDecrypt(encrypted, iv, key); err != nil {
		t.Error(err.Error())
		panic("Decryption failed")
	}
	if plaintext, err = RemovePadding(plaintext); err != nil {
		t.Error(err.Error())
		panic("Remove padding failed")
	}
	if data, err = RemovePadding(data); err != nil {
		t.Error(err.Error())
		panic("Remove padding failed")
	}
	if string(plaintext) != string(data) {
		t.Error(err.Error())
		panic("Doesn't match")
	}
}
