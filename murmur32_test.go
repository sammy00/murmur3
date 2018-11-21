package murmur3_test

import (
	"testing"

	"github.com/sammy00/murmur3"
	//	mm3 "github.com/spaolacci/murmur3"
)

func TestSumUint32(t *testing.T) {
	testCases := []struct {
		data   []byte
		seed   uint32
		expect uint32
	}{
		{
			seed:   0x9747b28c,
			data:   []byte("The quick brown fox jumps over the lazy dog"),
			expect: 0x2fa826cd,
		},
		{
			seed:   0x9747b28c,
			data:   []byte("The quick brown fox jumps over the lazy cog"),
			expect: 0x8d9c9ed4,
		},
		{
			seed:   0x9747b28c,
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"),
			expect: 0x637b8cf5,
		},
		{
			seed:   0x9747b28c,
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY COG"),
			expect: 0x21b8f236,
		},
		{
			seed:   0x9747b28c,
			data:   []byte("the quick brown fox jumps over the lazy dog"),
			expect: 0xedc02f1c,
		},
		{
			seed:   0x9747b28c,
			data:   []byte("the quick brown fox jumps over the lazy cog"),
			expect: 0x9781572b,
		},
		{
			seed:   0x0,
			data:   []byte("The quick brown fox jumps over the lazy dog"),
			expect: 0x2e4ff723,
		},
		{
			seed:   0x0,
			data:   []byte("The quick brown fox jumps over the lazy cog"),
			expect: 0xf08200fc,
		},
		{
			seed:   0x0,
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"),
			expect: 0x2c1bb0ae,
		},
		{
			seed:   0x0,
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY COG"),
			expect: 0x21b45485,
		},
		{
			seed:   0x0,
			data:   []byte("the quick brown fox jumps over the lazy dog"),
			expect: 0x2de62ff,
		},
		{
			seed:   0x0,
			data:   []byte("the quick brown fox jumps over the lazy cog"),
			expect: 0xbd613341,
		},
		{
			seed:   0xc58f1a7b,
			data:   []byte("The quick brown fox jumps over the lazy dog"),
			expect: 0x75b422d6,
		},
		{
			seed:   0xc58f1a7b,
			data:   []byte("The quick brown fox jumps over the lazy cog"),
			expect: 0x27757527,
		},
		{
			seed:   0xc58f1a7b,
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"),
			expect: 0xf24960e0,
		},
		{
			seed:   0xc58f1a7b,
			data:   []byte("THE QUICK BROWN FOX JUMPS OVER THE LAZY COG"),
			expect: 0x1ea370fc,
		},
		{
			seed:   0xc58f1a7b,
			data:   []byte("the quick brown fox jumps over the lazy dog"),
			expect: 0x14561460,
		},
		{
			seed:   0xc58f1a7b,
			data:   []byte("the quick brown fox jumps over the lazy cog"),
			expect: 0xb6d292ff,
		},
	}

	for i, c := range testCases {
		if got := murmur3.SumUint32(c.data, c.seed); got != c.expect {
			t.Fatalf("#%d invalid hash: got 0x%x, expect 0x%x", i, got, c.expect)
		}
	}
}
