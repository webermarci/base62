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
BenchmarkEncode/10^0_base62-8         	55492291	        21.08 ns/op	      64 B/op	       1 allocs/op
BenchmarkEncode/10^0_bigint-8         	15744423	        75.77 ns/op	      56 B/op	       4 allocs/op
BenchmarkEncode/10^5_base62-8         	32371807	        36.48 ns/op	      67 B/op	       2 allocs/op
BenchmarkEncode/10^5_bigint-8         	14995328	        79.29 ns/op	      56 B/op	       4 allocs/op
BenchmarkEncode/10^10_base62-8        	27377154	        43.05 ns/op	      72 B/op	       2 allocs/op
BenchmarkEncode/10^10_bigint-8        	13851474	        87.00 ns/op	      56 B/op	       4 allocs/op
BenchmarkEncode/10^15_base62-8        	22859700	        52.01 ns/op	      80 B/op	       2 allocs/op
BenchmarkEncode/10^15_bigint-8        	12992707	        92.74 ns/op	      64 B/op	       4 allocs/op
BenchmarkDecode/10^0_base62-8         	289116829	         4.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecode/10^0_bigint-8         	20567358	        57.53 ns/op	      40 B/op	       2 allocs/op
BenchmarkDecode/10^5_base62-8         	75475058	        15.78 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecode/10^5_bigint-8         	17250704	        69.33 ns/op	      40 B/op	       2 allocs/op
BenchmarkDecode/10^10_base62-8        	37104078	        31.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecode/10^10_bigint-8        	13986447	        85.01 ns/op	      40 B/op	       2 allocs/op
BenchmarkDecode/10^15_base62-8        	23032647	        51.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecode/10^15_bigint-8        	11808858	       100.70 ns/op	      40 B/op	       2 allocs/op
PASS
ok  	github.com/webermarci/base62	20.629s
```