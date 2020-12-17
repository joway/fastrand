# fastrand

## Background

Golang using a single [globalRand](https://github.com/golang/go/blob/master/src/math/rand/rand.go#L293) object between
multiple goroutines, which cause a race condition. But if we use `rand.New()` directly, it just creates a
non-thread-safe [rngSource]() object.

Fastrand wrap the `rngSource` struct into sync.Pool, and make it could run faster about 8 times than `globalRand` and
keep it thread-safe.

However, if you no need to share a single rand object within multiple goroutines, you could create a private rand.Rand
object in each goroutine.

## Usage

```go
rd := fastrand.New()
rd.Intn(100)
```

## Benchmark

```text
BenchmarkStandardRand
BenchmarkStandardRand-8                         60870416                19.1 ns/op             0 B/op          0 allocs/op
BenchmarkFastRand
BenchmarkFastRand-8                             100000000               10.7 ns/op             0 B/op          0 allocs/op
BenchmarkUnsafetyFastRand
BenchmarkUnsafetyFastRand-8                     100000000               10.7 ns/op             0 B/op          0 allocs/op
BenchmarkStandardRandWithConcurrent
BenchmarkStandardRandWithConcurrent-8           18058663                67.8 ns/op             0 B/op          0 allocs/op
BenchmarkFastRandWithConcurrent
BenchmarkFastRandWithConcurrent-8               132542940                8.79 ns/op            0 B/op          0 allocs/op
```
