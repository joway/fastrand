package fastrand

import (
	"math/rand"
	"testing"
)

func BenchmarkStandardRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Intn(1000)
	}
}

func BenchmarkFastRand(b *testing.B) {
	rd := NewUnsafety()
	for i := 0; i < b.N; i++ {
		rd.Intn(1000)
	}
}

func BenchmarkUnsafetyFastRand(b *testing.B) {
	rd := NewUnsafety()
	for i := 0; i < b.N; i++ {
		rd.Intn(1000)
	}
}

func BenchmarkStandardRandWithConcurrent(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.Intn(1000)
		}
	})
}

func BenchmarkFastRandWithConcurrent(b *testing.B) {
	rd := New()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rd.Intn(1000)
		}
	})
}
