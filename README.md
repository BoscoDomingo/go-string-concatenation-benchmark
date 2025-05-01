# go-string-concatenation-benchmark
A quick benchmark to test the 3 most common ways to concatenate strings (`+`, `fmt.Sprintf` and `strings.Builder`).

To test on your machine, run:
```sh
go test -bench=. -benchmem
```

--- 

# Latest results
## 2025-05-01
```sh
$ go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: scratchpad-go
cpu: AMD Ryzen 5 5600X 6-Core Processor
BenchmarkConcatStrings_Plus-12                          62491845                19.56 ns/op            8 B/op          1 allocs/op
BenchmarkConcatStrings_Sprintf-12                       11922879                91.66 ns/op           40 B/op          3 allocs/op
BenchmarkConcatStrings_StringsBuilder-12                75532712                15.50 ns/op            8 B/op          1 allocs/op
BenchmarkConcatStringAndInt_Plus-12                     54548280                19.90 ns/op            5 B/op          1 allocs/op
BenchmarkConcatStringAndInt_Sprintf-12                  13095879                92.15 ns/op           21 B/op          2 allocs/op
BenchmarkConcatStringAndInt_StringsBuilder-12           64315933                18.32 ns/op            8 B/op          1 allocs/op
```

`Sprintf` is the slowest and takes the most memory.
`Strings.Builder` is the fastest yet probably the most cumbersome.
The `+` operator is the second fastest and tied for least memory, so it's probably the best option.
