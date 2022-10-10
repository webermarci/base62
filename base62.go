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

var (
	multipliers []uint64 = []uint64{
		1,
		62,
		3844,
		238328,
		14776336,
		916132832,
		56800235584,
		3521614606208,
		218340105584896,
		13537086546263552,
		839299365868340224,
	}
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
	length := len(encoded)

	for i := 0; i < length; i++ {
		index := strings.IndexByte(charset, encoded[i])
		if index == -1 {
			return 0, fmt.Errorf("character not found: %s", string(encoded[i]))
		}
		number += uint64(index) * multipliers[length-i-1]
	}
	return number, nil
}
