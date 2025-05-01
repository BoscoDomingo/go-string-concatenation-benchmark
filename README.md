# go-string-concatenation-benchmark
A quick benchmark to test the 4 most common ways to concatenate strings (`+`, `fmt.Sprintf`, `strings.Builder` and `strings.Join`).

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
BenchmarkConcatStrings_Plus-12                   	51331621	        21.97 ns/op	       8 B/op	       1 allocs/op
BenchmarkConcatStrings_Sprintf-12                	11596424	       102.9 ns/op	      40 B/op	       3 allocs/op
BenchmarkConcatStrings_StringsBuilder-12         	64022565	        18.46 ns/op	       8 B/op	       1 allocs/op
BenchmarkConcatStrings_Join-12                   	42522718	        27.70 ns/op	       8 B/op	       1 allocs/op
BenchmarkConcatStringAndInt_Plus-12              	51975021	        22.21 ns/op	       5 B/op	       1 allocs/op
BenchmarkConcatStringAndInt_Sprintf-12           	12212780	        98.96 ns/op	      21 B/op	       2 allocs/op
BenchmarkConcatStringAndInt_StringsBuilder-12    	56754538	        20.43 ns/op	       8 B/op	       1 allocs/op
BenchmarkConcatStringAndInt_Join-12              	38538142	        28.96 ns/op	       8 B/op	       1 allocs/op
```

The `+` operator is the second fastest and tied for least memory while also being the most convenient, so it's probably the best option in the majority of cases.
`Sprintf` is the slowest and takes the most memory. Don't use it.
`strings.Builder` is the fastest yet probably the most cumbersome.
`strings.Join` is a great option for arrays of strings.
