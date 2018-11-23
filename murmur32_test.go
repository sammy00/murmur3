package murmur3_test

import (
	"testing"

	"github.com/sammy00/murmur3"
)

func TestSum32(t *testing.T) {
	testCases := []struct {
		data   []byte
		seed   uint32
		expect [4]byte
	}{
		{
			data:   []byte("The quick brown fox jumps over the lazy dog"),
			seed:   0x9747b28c,
			expect: [4]byte{0xcd, 0x26, 0xa8, 0x2f},
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy cog"),
			seed:   0x9747b28c,
			expect: [4]byte{0xd4, 0x9e, 0x9c, 0x8d},
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"),
			seed:   0x9747b28c,
			expect: [4]byte{0xf5, 0x8c, 0x7b, 0x63},
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY COG"),
			seed:   0x9747b28c,
			expect: [4]byte{0x36, 0xf2, 0xb8, 0x21},
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy dog"),
			seed:   0x9747b28c,
			expect: [4]byte{0x1c, 0x2f, 0xc0, 0xed},
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog"),
			seed:   0x9747b28c,
			expect: [4]byte{0x2b, 0x57, 0x81, 0x97},
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy dog"),
			seed:   0x00000000,
			expect: [4]byte{0x23, 0xf7, 0x4f, 0x2e},
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy cog"),
			seed:   0x00000000,
			expect: [4]byte{0xfc, 0x00, 0x82, 0xf0},
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"),
			seed:   0x00000000,
			expect: [4]byte{0xae, 0xb0, 0x1b, 0x2c},
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY COG"),
			seed:   0x00000000,
			expect: [4]byte{0x85, 0x54, 0xb4, 0x21},
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy dog"),
			seed:   0x00000000,
			expect: [4]byte{0xff, 0x62, 0xde, 0x02},
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog"),
			seed:   0x00000000,
			expect: [4]byte{0x41, 0x33, 0x61, 0xbd},
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy dog"),
			seed:   0xc58f1a7b,
			expect: [4]byte{0xd6, 0x22, 0xb4, 0x75},
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy cog"),
			seed:   0xc58f1a7b,
			expect: [4]byte{0x27, 0x75, 0x75, 0x27},
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"),
			seed:   0xc58f1a7b,
			expect: [4]byte{0xe0, 0x60, 0x49, 0xf2},
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY COG"),
			seed:   0xc58f1a7b,
			expect: [4]byte{0xfc, 0x70, 0xa3, 0x1e},
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy dog"),
			seed:   0xc58f1a7b,
			expect: [4]byte{0x60, 0x14, 0x56, 0x14},
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog"),
			seed:   0xc58f1a7b,
			expect: [4]byte{0xff, 0x92, 0xd2, 0xb6},
		},
	}

	for i, c := range testCases {
		if got := murmur3.Sum32(c.data, c.seed); got != c.expect {
			t.Fatalf("#%d invalid hash: got 0x%x, expect 0x%x", i, got, c.expect)
		}
	}
}

func TestSumUint32(t *testing.T) {
	testCases := []struct {
		data   []byte
		seed   uint32
		expect uint32
	}{
		{
			data:   []byte("The quick brown fox jumps over the lazy dog"),
			seed:   0x9747b28c,
			expect: 0x2fa826cd,
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy cog"),
			seed:   0x9747b28c,
			expect: 0x8d9c9ed4,
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"),
			seed:   0x9747b28c,
			expect: 0x637b8cf5,
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY COG"),
			seed:   0x9747b28c,
			expect: 0x21b8f236,
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy dog"),
			seed:   0x9747b28c,
			expect: 0xedc02f1c,
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog"),
			seed:   0x9747b28c,
			expect: 0x9781572b,
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy dog"),
			seed:   0x00000000,
			expect: 0x2e4ff723,
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy cog"),
			seed:   0x00000000,
			expect: 0xf08200fc,
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"),
			seed:   0x00000000,
			expect: 0x2c1bb0ae,
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY COG"),
			seed:   0x00000000,
			expect: 0x21b45485,
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy dog"),
			seed:   0x00000000,
			expect: 0x02de62ff,
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog"),
			seed:   0x00000000,
			expect: 0xbd613341,
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy dog"),
			seed:   0xc58f1a7b,
			expect: 0x75b422d6,
		},
		{
			data:   []byte("The quick brown fox jumps over the lazy cog"),
			seed:   0xc58f1a7b,
			expect: 0x27757527,
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"),
			seed:   0xc58f1a7b,
			expect: 0xf24960e0,
		},
		{
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY COG"),
			seed:   0xc58f1a7b,
			expect: 0x1ea370fc,
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy dog"),
			seed:   0xc58f1a7b,
			expect: 0x14561460,
		},
		{
			data:   []byte("the quick brown fox jumps over the lazy cog"),
			seed:   0xc58f1a7b,
			expect: 0xb6d292ff,
		},
	}

	for i, c := range testCases {
		if got := murmur3.SumUint32(c.data, c.seed); got != c.expect {
			t.Fatalf("#%d invalid hash: got 0x%x, expect 0x%x", i, got, c.expect)
		}
	}
}
