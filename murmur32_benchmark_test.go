package murmur3_test

import (
	"strconv"
	"testing"

	"github.com/sammyne/murmur3"
)

func BenchmarkSum32(b *testing.B) {
	data := make([]byte, 8196)
	const Seed = 0x00000000

	for n := 1; n <= cap(data); n *= 2 {
		b.Run(strconv.Itoa(n), func(sb *testing.B) {
			sb.SetBytes(int64(n))
			sb.ResetTimer()

			for i := 0; i < sb.N; i++ {
				murmur3.Sum32(data[:n], Seed)
			}
		})
	}
}
