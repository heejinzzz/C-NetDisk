package handler

import (
	"encoding/base64"
	"log"
)

var (
	characterSet = []byte{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '_', '!',
	}
	encoding = base64.NewEncoding(string(characterSet))
)

func base64Encode(s string) string {
	res := make([]byte, encoding.EncodedLen(len(s)))
	encoding.Encode(res, []byte(s))
	return string(res)
}

func base64Decode(s string) string {
	res := make([]byte, encoding.DecodedLen(len(s)))
	_, err := encoding.Decode(res, []byte(s))
	if err != nil {
		log.Fatalln(err)
	}
	return string(res)
}
