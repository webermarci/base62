package base62

import (
	"fmt"
	"math"
	"math/big"
	"testing"
)

func TestEncoding(t *testing.T) {
	for i := 0; i <= 1000000; i++ {
		encoded := Encode(uint64(i))
		decoded, err := Decode(encoded)
		if err != nil {
			t.Fatal(err)
		}

		if uint64(i) != decoded {
			t.Fatalf("%d != %d", int64(i), decoded)
		}
	}
}

func TestInvalidCharacterDecoding(t *testing.T) {
	_, err := Decode("...")
	if err == nil {
		t.Fatal("expected error")
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i <= 15; i += 5 {
		number := uint64(math.Pow10(i))

		b.Run(fmt.Sprintf("10^%d base62", i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Encode(number)
			}
		})

		b.Run(fmt.Sprintf("10^%d bigint", i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				big.NewInt(int64(number)).Text(62)
			}
		})
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i <= 15; i += 5 {
		number := uint64(math.Pow10(i))
		encoded := Encode(number)

		b.Run(fmt.Sprintf("10^%d base62", i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Decode(encoded)
			}
		})

		b.Run(fmt.Sprintf("10^%d bigint", i), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var bigInt big.Int
				bigInt.SetString(encoded, 62)
			}
		})
	}
}
