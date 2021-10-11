package popcount

import (
	"testing"
)
import "math/rand"
import "time"

func BenchmarkPopCount(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		PopCount(rand.Uint64())
	}
}

func BenchmarkPopCountV2(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		PopCountV2(rand.Uint64())
	}
}

func BenchmarkPopCountV3(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		PopCountV3(rand.Uint64())
	}
}

func BenchmarkPopCountV4(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		PopCountV4(rand.Uint64())
	}
}
