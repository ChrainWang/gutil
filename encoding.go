package gutil

import (
	"crypto/md5"
	"encoding/base64"
)

func Base64Encode(source []byte) (dst []byte) {
	encodeLength := base64.StdEncoding.EncodedLen(len(source))
	dst = make([]byte, encodeLength)
	base64.StdEncoding.Encode(dst, source)
	return
}

func Base64Decode(source []byte) (dst []byte, err error) {
	length := base64.StdEncoding.EncodedLen(len(source))
	dst = make([]byte, length)
	var n int
	if n, err = base64.StdEncoding.Decode(dst, source); err == nil {
		dst = dst[:n]
	}
	return
}

func Md5Encode(source []byte) (dst []byte) {
	hasher := md5.New()
	hasher.Write(source)
	dst = hasher.Sum(nil)
	return
}
