# Go Resty v2 vs v3 Benchmark: Marshal and Unmarshal JSON

## 💻 System Information

| Component | Detail       |
| --------- | ------------ |
| OS        | macOS 15.4.1 |
| CPU       | Apple M4     |
| RAM       | 16 GB        |

## 📋 Prerequisites

- Go 1.24.5

```sh
make install
```

## 📊 Results

- `make bench`

```go
goos: darwin
goarch: arm64
pkg: go-resty-v2-v3
cpu: Apple M4
BenchmarkMarshal/Small/RestyV2_StdMarshaler-10             26767             42638 ns/op           26215 B/op        103 allocs/op
BenchmarkMarshal/Small/RestyV2_GoccyMarshaler-10           29199             40761 ns/op           28193 B/op        103 allocs/op
BenchmarkMarshal/Small/RestyV3_StdEncoder-10               31646             37768 ns/op            9519 B/op         99 allocs/op
BenchmarkMarshal/Small/RestyV3_GoccyEncoder-10             33417             35778 ns/op            9595 B/op         99 allocs/op
BenchmarkMarshal/Medium/RestyV2_StdMarshaler-10             9162            119470 ns/op          272163 B/op        121 allocs/op
BenchmarkMarshal/Medium/RestyV2_GoccyMarshaler-10           9664            115041 ns/op          323241 B/op        124 allocs/op
BenchmarkMarshal/Medium/RestyV3_StdEncoder-10              13431             89011 ns/op           51098 B/op        101 allocs/op
BenchmarkMarshal/Medium/RestyV3_GoccyEncoder-10            16033             74890 ns/op           73972 B/op        104 allocs/op
BenchmarkMarshal/Large/RestyV2_StdMarshaler-10              1508            954893 ns/op         1933452 B/op        225 allocs/op
BenchmarkMarshal/Large/RestyV2_GoccyMarshaler-10            1390            894173 ns/op         2701207 B/op        229 allocs/op
BenchmarkMarshal/Large/RestyV3_StdEncoder-10                1470            923219 ns/op          372757 B/op        198 allocs/op
BenchmarkMarshal/Large/RestyV3_GoccyEncoder-10              1809            637595 ns/op         1083730 B/op        213 allocs/op
BenchmarkUnmarshal/Small/RestyV2_StdUnmarshaler-10         15186             78739 ns/op           38766 B/op        394 allocs/op
BenchmarkUnmarshal/Small/RestyV2_GoccyUnmarshaler-10       18799             63565 ns/op           47928 B/op        251 allocs/op
BenchmarkUnmarshal/Small/RestyV3_StdDecoder-10             15316             78219 ns/op           33595 B/op        394 allocs/op
BenchmarkUnmarshal/Small/RestyV3_GoccyDecoder-10           20695             58061 ns/op           24735 B/op        245 allocs/op
BenchmarkUnmarshal/Medium/RestyV2_StdUnmarshaler-10         2434            493837 ns/op          405788 B/op       2942 allocs/op
BenchmarkUnmarshal/Medium/RestyV2_GoccyUnmarshaler-10       3673            322885 ns/op          426182 B/op       1540 allocs/op
BenchmarkUnmarshal/Medium/RestyV3_StdDecoder-10             2564            455524 ns/op          258235 B/op       2922 allocs/op
BenchmarkUnmarshal/Medium/RestyV3_GoccyDecoder-10           2247            454340 ns/op          223367 B/op       1619 allocs/op
BenchmarkUnmarshal/Large/RestyV2_StdUnmarshaler-10           298           4007366 ns/op         4592401 B/op      38605 allocs/op
BenchmarkUnmarshal/Large/RestyV2_GoccyUnmarshaler-10         571           2090069 ns/op         5178090 B/op      24608 allocs/op
BenchmarkUnmarshal/Large/RestyV3_StdDecoder-10               309           3836035 ns/op         2870352 B/op      38575 allocs/op
BenchmarkUnmarshal/Large/RestyV3_GoccyDecoder-10             552           2159410 ns/op         2105529 B/op      24657 allocs/op
```
