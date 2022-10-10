package base62

import (
	"fmt"
	"math"
	"math/big"
	"testing"
	"time"
)

func TestEncoding(t *testing.T) {
	testData := []uint64{
		0,
		1,
		42,
		69,
		uint64(time.Now().Unix()),
		uint64(time.Now().UnixMilli()),
		uint64(time.Now().UnixMicro()),
		uint64(time.Now().UnixNano()),
	}

	for _, number := range testData {
		encoded := Encode(number)
		decoded, err := Decode(encoded)
		if err != nil {
			t.Fatal(err)
		}

		if number != decoded {
			t.Fatalf("%d != %d", number, decoded)
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
