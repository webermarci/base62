package base62

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	base    = 62
)

func Encode(number uint64) string {
	var buf bytes.Buffer
	encode(number, &buf)
	return buf.String()
}

func encode(number uint64, buf *bytes.Buffer) {
	if number/base != 0 {
		encode(number/base, buf)
	}
	buf.WriteByte(charset[number%base])
}

func Decode(encoded string) (uint64, error) {
	var number uint64 = 0
	var multiplier uint64
	length := len(encoded)

	for i := 0; i < length; i++ {
		multiplier = 1
		for j := 0; j < length-i-1; j++ {
			multiplier *= base
		}
		index := strings.IndexByte(charset, encoded[i])
		if index == -1 {
			return 0, fmt.Errorf("character not found: %s", string(encoded[i]))
		}
		number += uint64(index) * multiplier
	}
	return number, nil
}
