# Base 62 Encoder

## Usage
```go
import "github.com/webermarci/base62"

var number uint64 = 1665401943785138000
encoded := base62.Encode(number)
// 1z1Yd99Ld0i

decoded, err := base62.Decode(encoded)
// 1665401943785138000
```

## Performance
It's a bit faster than the standard `math/big.Int` implementation.

```
goos: darwin
goarch: arm64
pkg: github.com/webermarci/base62
BenchmarkEncode/10^0_base62-8         	56710610	        20.99 ns/op	      64 B/op	       1 allocs/op
BenchmarkEncode/10^0_bigint-8         	15800838	        75.11 ns/op	      56 B/op	       4 allocs/op
BenchmarkEncode/10^5_base62-8         	32987422	        35.64 ns/op	      67 B/op	       2 allocs/op
BenchmarkEncode/10^5_bigint-8         	15108235	        78.66 ns/op	      56 B/op	       4 allocs/op
BenchmarkEncode/10^10_base62-8        	27390901	        42.71 ns/op	      72 B/op	       2 allocs/op
BenchmarkEncode/10^10_bigint-8        	13797662	        86.18 ns/op	      56 B/op	       4 allocs/op
BenchmarkEncode/10^15_base62-8        	23152762	        51.31 ns/op	      80 B/op	       2 allocs/op
BenchmarkEncode/10^15_bigint-8        	13191324	        91.99 ns/op	      64 B/op	       4 allocs/op
BenchmarkDecode/10^0_base62-8         	292363821	         4.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecode/10^0_bigint-8         	20652962	        56.90 ns/op	      40 B/op	       2 allocs/op
BenchmarkDecode/10^5_base62-8         	89711798	        12.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecode/10^5_bigint-8         	17419663	        68.31 ns/op	      40 B/op	       2 allocs/op
BenchmarkDecode/10^10_base62-8        	49069218	        24.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecode/10^10_bigint-8        	14067397	        83.87 ns/op	      40 B/op	       2 allocs/op
BenchmarkDecode/10^15_base62-8        	35270646	        33.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecode/10^15_bigint-8        	12042242	        99.09 ns/op	      40 B/op	       2 allocs/op
PASS
ok  	github.com/webermarci/base62	21.386s
```